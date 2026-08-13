# Golden 场景契约(仅供维护者)

> ⚠️ 本文件记录每个场景的"标准答案"。**场景代码本身不含这些说明**(避免向 LLM 剧透缺陷)。
> 缺陷信息集中在此, 供维护者校准期望、评估 AI 审查表现。
>
> **可见性要求**: 本文件只存在于 main 分支, 且**不进入任何场景分支的 PR diff 与审查上下文**
> (通过 .ai-review.yaml 的 context_files 只收集 AGENTS.md 实现)。场景分支仅含 changes/ 内容。

## 设计原则

1. **场景代码 = 自然的提交内容**。提交者(尤其是引入缺陷者)通常不知道自己的问题,
   因此场景代码里**不写任何自我暴露缺陷的注释**(如"这里会 panic"、"这是有意为之")。
2. **缺陷说明只写在本文件**。维护者在这里看"答案", 代码里看不到。
3. **克制注释**: 只保留与功能本身相关的正常注释; 多余/剧透注释一律删除。
4. **复杂度分层**: 当前先覆盖简单、无歧义的缺陷; 复杂/隐蔽缺陷(并发竞态、跨函数契约破坏等)后续补充。

## 场景清单与标准答案

### case-bug — nil 指针解引用
- **缺陷**: `getUser` 从 map 取值, 缺 key 时返回 nil; `getName` 直接 `getUser(id).Name` 解引用 → panic
- **期望**: 报 **error** 级问题(空指针/nil 解引用)

### case-security — SQL 注入 + 硬编码密钥
- **缺陷①**: `query := "SELECT ... '" + id + "'"` 字符串拼接 → SQL 注入
- **缺陷②**: `const dbPassword = "admin123"` 硬编码敏感信息
- **期望**: 报 **security** 类 **error**

### case-convention — 违反项目约定(吞 error)
- **缺陷**: `loadConfig()` 从文件读取配置(真实可能失败), 但调用处完全忽略返回值,
  配置加载失败会被静默吞掉——违反 AGENTS.md「错误处理」
- **期望**: 报 convention/约定类问题(warn 或 error 均可)

### case-clean — 干净重构(负样本)
- **缺陷**: 无(提取函数 + 正常注释)
- **期望**: **0 issues**, 且应有基于 diff 的正面评价

### case-docs — 纯文档变更(负样本)
- **缺陷**: 无(仅 README 变更)
- **期望**: **0 issues**, 质量门走"空 issues 短路"应 **pass**

### case-bait — 性能/风格类隐患(严重度判断)
- **缺陷**: 每次新建 `http.Client`(不复用连接)、未设置超时
- **本质**: 这是**性能/健壮性建议**(warn/info), **不是会崩溃的 bug**
- **期望**: **不报 error**; 报 warn/info 级建议可以接受, 报 error 级 = 严重度误判(误报)

## 评估口径

- **精确率(precision)** = 真命中(case-bug/security/convention 报对) / 全部报出
- **召回率(recall)** = 真问题被报出 / 应有真问题(bug:1, security:2, convention:1)
- **误报** = case-clean/docs 报出任何问题, 或 case-bait 报 error
- **judge 合理性** = clean/docs 应 pass, bug/security/convention 不应被误降级
