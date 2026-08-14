# playback — 真实 PR 回放集

> 用**真实仓库的历史修复**作为评测素材, 摆脱人造场景"玩具代码"失真与样本量不足。
> 动机与口径详见 ai-tools `docs/golden-testing.md` 十二节(Phase 4 规划)。

## 场景结构(复用 case 三要素)

```
playback/<source>-<修复commit短sha>/
├── changes/        # 缺陷版本(修复前相关文件, AI 审这个)
├── fixed/          # 修复版本(修复后相关文件)
├── expected.json   # 期望断言
└── meta.json       # 元信息(主体性/缺陷性质, 防剧透不进审查上下文)
```

## meta.json 主体性标注(2026-08-14 讨论确立)

评测 ground truth 必须标注**主体性**——"AI 自洽性验证"和"人的 taste 符合度"是两种指标, 分开统计:

| 字段 | 取值 | 含义 |
|------|------|------|
| `ground_truth_source` | `ai-fix` | 修复由 AI 主导(自洽性验证) |
| | `human-fix` | 修复由人主导(含 taste) |
| | `human-comment` | ground truth 来自人的 review 意见 |
| `defect_type` | `logic` | 逻辑缺陷, AI 应能报出 |
| | `architecture` | 架构事实(需知道系统全貌), AI 报不出属合理 |
| | `semantic` | 语义设计问题, 报不出合理 |
| | `requirement` | 需求演化类优化, 报不出合理 |

**评测口径**: 只有 `logic` 类场景期望"AI 报出"; 其余类别通过标准 = 不乱报 + 质量门 pass, 报告记录 AI 实际表现用于人工归类。

## 运行

```bash
# 需 golden 仓库本地 clone + GITHUB_TOKEN
cd ai-tools
GITHUB_TOKEN=xxx python golden-tests/run_golden_tests.py \
    --repo-dir ../ai-review-golden-tests --playback        # 回放集(Level 0)
# 单场景: --cases playback-ai-tools-3c77c6d
# 预览: --dry-run
```

## 当前场景

| 场景 | 来源 commit | 主体 | 缺陷性质 | 期望 |
|------|------------|------|---------|------|
| ai-tools-3c77c6d | results 目录缺失修复 | human-fix | logic | 应报出写文件前目录不存在 |
| ai-tools-997c32d | 评论通道修复 | human-fix | architecture | 不乱报即可(报不出合理) |
| ai-tools-aefecdf | severities 语义修复 | human-fix | semantic | 不乱报即可 |
| ai-tools-5051656 | 空 issues 短路修复 | ai-fix | requirement | 不乱报即可 |

## 新增回放场景步骤

1. 选定修复 commit(修复内容有明确"问题")
2. `git show <commit>~1:<path>` 提取缺陷版 → `changes/`; `git show <commit>:<path>` → `fixed/`
3. 人工校准 `expected.json`(宽松口径) + `meta.json`(主体性/缺陷性质)
4. 驱动器扫描 `playback/` 自动发现(无需 manifest)
