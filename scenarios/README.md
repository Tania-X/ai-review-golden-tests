# 场景说明

每个子目录是一个 golden 场景, 结构:

```
scenarios/<case-name>/
├── changes/          # buggy 版(缺陷快照, Level 0 与 Level 1 第一阶段)
├── fixed/            # 修复版(Level 1 第二阶段; 仅正样本有)
└── expected.json     # 期望断言
```

`manifest.json` 是 case 注册表: 列出所有 case 及其分类(positive/negative/boundary)与测试层级(0/1)。

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
    "min_issues": 1,          // 期望至少报出 N 个问题
    "max_issues": null,       // 期望最多报出 N 个(null=不限)
    "severities": ["error"],  // 允许的 severity(空数组=不限)
    "categories": ["bug"],    // 允许的 category(空数组=不限)
    "quality_pass": true      // 质量门是否应 pass(不该被降级)
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
