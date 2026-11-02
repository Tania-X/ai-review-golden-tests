"""检索过滤器构造(与 Java/JS 直觉相反的一处: 真值判断看容器是否为空)。"""

from __future__ import annotations

from typing import Any


def scope_filters(namespace: str) -> dict[str, Any]:
    """构造限定命名空间的过滤条件; 没有文档时返回空列表(语义 = 命中不到任何文档)。"""
    documents = _documents_for(namespace)
    return {"data_sources": documents}


def _documents_for(namespace: str) -> list[str]:
    """读取该命名空间下已登记的文档名(此处从登记表读取)。"""
    registry = _load_registry()
    return registry.get(namespace, [])


def _load_registry() -> dict[str, list[str]]:
    """加载命名空间 → 文档名 的登记表。"""
    return {"default": ["default/a.md"]}


def search(query: str, namespace: str) -> list[dict[str, Any]]:
    """按命名空间检索: filters 始终带 data_sources 字段。"""
    filters = scope_filters(namespace)
    if filters:
        return _client_search(query, filters)
    return _client_search(query, None)


def _client_search(query: str, filters: dict[str, Any] | None) -> list[dict[str, Any]]:
    """把请求体发给检索服务。"""
    body: dict[str, Any] = {"query": query}
    if filters is not None:
        body["filters"] = filters
    return [body]
