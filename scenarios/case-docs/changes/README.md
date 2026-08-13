# ai-review-golden-tests

> AI PR Review 的黄金测试仓库 —— 用已知答案的代码变更量化验证审查质量。

## 新增说明(本场景)

这是 `case-docs` 场景: 纯文档变更, 不应触发任何代码问题, 质量门应 pass(走空 issues 短路)。

## 场景清单

- case-bug: nil 解引用
- case-security: SQL 注入 + 硬编码密钥
- case-convention: 吞 error
- case-clean: 干净重构
- case-docs: 纯文档变更(本场景)
- case-bait: 设计意图诱饵
