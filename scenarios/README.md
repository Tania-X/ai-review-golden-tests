# 场景说明

每个子目录是一个 golden 场景, 结构:

```
scenarios/<case-name>/
├── changes/          # buggy 版(缺陷快照, Level 0 与 Level 1 第一阶段)
├── fixed/            # 修复版(Level 1 第二阶段; 仅正样本有)
└── expected.json     # 期望断言
```

`manifest.json` 是 case 注册表: 列出所有 case 及其分类(positive/negative/boundary)、测试层级(0/1)、
**能力维度(capability)** 与 **语言(language)**。

> capability 是"审查能力"维度(语言无关, 报告按它聚合通过率):
> `bug / security / convention / severity / merge-locations / context / no-false-positive /
> language-semantics`
> language 是场景载体语言(go / python)。

## language-semantics(2026-09-13 新增)

针对**语言先验误报**: 正确代码 + 与其它语言(Java/JS)直觉相反的语言规则。
这组用例全部是 **negative**(期望不被报), 判据用 `forbid_severities: [3,4,5]`:
允许 1-2 级风格建议, 但把正确代码判成必修及以上即为误报。

| case | 与直觉相反的点 |
|------|----------------|
| case-semantics-truthiness | 含空列表的 dict 仍为真(`{"data_sources": []}` 不会被丢弃) |
| case-semantics-exception-mapping | 异常按类集中注册、沿 MRO 取最精确处理器, 路由内无需逐个 try/except |
| case-semantics-async-scope | async generator 必须用 asynccontextmanager 包裹; `async with` 内 await 不会提前释放资源 |
| case-semantics-dataclass | 普通 dataclass(含 frozen)**不校验**参数值, 非法值表现为查不到而非崩溃 |
| case-semantics-config-default | `x or default` 的 default 来自配置对象, 不是硬编码常量 |

`stack.languages` 已在仓库 `.ai-review.yaml` 声明为 `[go, python]`, 引擎会注入对应的语义清单。

## case 分类

| 类别 | case | 审查应拒绝? | 层级 |
|------|------|------------|------|
| positive | bug / security / convention | 应 refuse(报 error) | 0 + 1 |
| negative | clean / docs | 应 agree(无问题) | 0 |
| boundary | bait | 应 agree(可 warn/info, 不报 error) | 0 |

## expected.json schema

```jsonc
{
  "case": "case-bug",
  "description": "注入 nil 解引用 bug",
  "expect": {
    "min_issues": 1,              // 期望至少报出 N 个问题
    "max_issues": null,           // 期望最多报出 N 个(null=不限)
    "severities": ["error"],      // 必须命中其一(报出级别中至少一个在此集合内)
    "forbid_severities": ["error"], // 禁止报出的级别(如边界样本禁止 error)
    "categories": ["bug"],        // 允许的 category(空数组=不限; 当前断言器未校验, 待实现)
    "quality_pass": true          // 质量门是否应 pass(不该被降级)
  }
}
```

## Level 0 手工验证(单次审查)

```bash
git checkout -b test/case-bug
cp -r scenarios/case-bug/changes/* .
git add . && git commit -m "add user query"
git push origin test/case-bug
# 开 PR → 观察 AI review 表现 → 比对 expected.json
```

## Level 1 手工验证(修复闭环, 正样本)

```bash
# 接 Level 0: review 应 refuse 后
cp -r scenarios/case-bug/fixed/* .
git add . && git commit -m "fix review findings"
git push origin test/case-bug
# review 重跑(synchronize)→ 应 agree → 可 merge
```

完整方法论见 ai-tools 仓库 `docs/golden-testing.md`。
