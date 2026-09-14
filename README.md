# ai-review-golden-tests

> 记录 AI PR Review 黄金测试仓库的使用说明。

## 快速开始

1. 克隆本仓库
2. 参考 `scenarios/README.md` 了解场景结构
3. 按需创建场景分支, 观察 pr-review 的表现

## 依赖

- AI 审查由 GitHub Actions 自动触发(见 `.github/workflows/ai-review.yml`)
- 需要仓库 Secret: `DEEPSEEK_API_KEY`
