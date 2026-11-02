"""值对象与登记表查询(普通 dataclass 不校验参数值, 非法值只表现为查不到)。"""

from __future__ import annotations

from dataclasses import dataclass, field


@dataclass(frozen=True, slots=True)
class TenantId:
    """租户标识(不做取值校验: 非空字符串即可构造)。"""

    value: str

    def __str__(self) -> str:
        return self.value


@dataclass(slots=True)
class Document:
    """文档登记行。"""

    name: str
    display_name: str
    tags: list[str] = field(default_factory=list)


_TENANTS: dict[str, list[Document]] = {
    "default": [Document(name="default/a.md", display_name="A")],
}


def lookup(tenant: TenantId) -> list[Document] | None:
    """按租户查文档; 查不到返回 None(由调用层翻译成 404)。"""
    return _TENANTS.get(str(tenant))


def documents_for(raw_tenant: str) -> list[Document]:
    """接口层入口: 非法/未知租户与"不存在的租户"走同一条路(404), 不会抛异常。"""
    found = lookup(TenantId(raw_tenant))
    if found is None:
        return []
    return found
