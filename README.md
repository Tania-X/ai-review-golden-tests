# ai-review-golden-tests

> AI PR Review 的**黄金测试仓库** —— 用一组"已知答案"的代码变更，量化验证 ai-tools/pr-review 的审查质量（精确率 / 召回率 / 打分合理性）。

## 这个项目是做什么的

ai-tools 的 pr-review 是一个 LLM 审查 PR 的 GitHub Action。它的质量（会不会误报、会不会漏报、judge 打分准不准）**无法靠单元测试验证**——因为 LLM 输出是非确定性的，只有跑在真实 GitHub 环境里、对真实 diff 审查，才能看到真实表现。

本仓库就是为此而生：**预置一批"标准答案已知"的代码变更场景**，通过自动化流程让 pr-review 逐个审查，再比对"实际输出 vs 期望输出"，得到可重复、可量化的质量报告。

## 工作原理

```
golden 场景(已知 bug / 干净代码 / 文档变更 / 误报诱饵)
        │
        ▼  驱动器: 建分支 → 变更 → push → 开 PR
        │  (pr-review 自动跑, 异步)
        ▼  轮询 check-run / 评论 → 解析 issues
        │
        ▼  断言 expected.json vs 实际输出
        │
        ▼  量化报告: 精确率 / 召回率 / judge 合理性 / 场景矩阵
```

每次修改 ai-tools 的审查 prompt 或打分逻辑后，跑一遍全量场景，就能看到质量是提升还是退化。

## 目录结构

```
.
├── .github/workflows/ai-review.yml   # 接入 pr-review(引用 ai-tools@main)
├── .ai-review.yaml                    # 审查配置(默认,质量门开启)
├── AGENTS.md                          # 项目约定(case-convention 场景的依据)
├── src/main.go                        # 基线代码(干净版, main 分支)
└── scenarios/                         # 场景定义(不参与基线, 供驱动器逐场景变更)
    ├── README.md                      # 场景说明
    ├── case-bug/                      # 已知 bug(nil 解引用)
    ├── case-security/                 # 安全问题(SQL 注入/硬编码密钥)
    ├── case-convention/               # 违反项目约定(吞 error)
    ├── case-clean/                    # 干净重构(应无问题)
    ├── case-docs/                     # 纯文档变更(应走质量门短路)
    └── case-bait/                     # 误报诱饵(设计意图, 不应确定性报错)
```

每个场景目录含 `changes/`(该场景 PR 应引入的文件, 镜像仓库根路径)和 `expected.json`(期望断言)。

## 场景矩阵

| 场景 | 变更 | 期望(核心断言) |
|------|------|----------------|
| `case-bug` | nil 指针解引用 | 报 error 级问题 |
| `case-security` | SQL 拼接 + 硬编码密钥 | 报 security 类问题 |
| `case-convention` | 吞掉 error(违反 AGENTS.md) | 报 convention/约定类问题 |
| `case-clean` | 干净重构 | **0 issues + 正面评价** |
| `case-docs` | 纯文档变更 | 0 issues + 质量门 pass |
| `case-bait` | 看起来像问题、实为设计意图 | **不报确定性 error**(可标 needs_review) |

## 量化指标

跑完一套场景可得出:

- **精确率(precision)** = 真命中数 / 全部报出数 —— `case-bug`/`case-security` 报对的比例, 反映"误报"水平
- **召回率(recall)** = 真问题被报出数 / 应有真问题数 —— 反映"漏报"水平
- **误报** = `case-clean`/`case-bait` 不该报却报了
- **judge 合理性** = 干净/文档场景是否 pass、有 bug 场景是否未被误降级

## 当前状态

- **Phase 1(当前)**: 场景定义 + 期望已就位; 可手工逐个场景开 PR 观察 pr-review 表现
- **Phase 2(规划)**: 自动化驱动器 `run_golden_tests.py`(建分支→开 PR→轮询→断言→报告)
