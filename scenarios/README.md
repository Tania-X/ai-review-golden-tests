# 场景说明

每个子目录是一个 golden 场景, 结构:

```
scenarios/<case-name>/
├── changes/          # 该场景 PR 应引入的文件(镜像仓库根路径)
│   └── src/main.go   # 例: 替换基线的 main.go
└── expected.json     # 期望断言
```

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

## Phase 1 手工验证方式

对单个场景, 手工执行:

```bash
git checkout -b test/case-bug
cp -r scenarios/case-bug/changes/* .
git add . && git commit -m "test: case-bug"
git push origin test/case-bug
# 在 GitHub 开 PR → 观察 AI review 表现 → 比对 expected.json
```

Phase 2 会把这些步骤自动化。
