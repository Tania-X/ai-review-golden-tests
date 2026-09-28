"""按类型集中分发领域异常(处理器注册在框架层, 调用点不再逐个捕获)。"""

from __future__ import annotations

from typing import Any

HANDLERS: dict[type[Exception], int] = {}


def handles(exc_type: type[Exception], status: int) -> None:
    """注册某类异常对应的 HTTP 状态码。"""
    HANDLERS[exc_type] = status


def status_for(exc: BaseException) -> int:
    """沿 MRO 找最精确的处理器; 找不到时按 500 处理。"""
    for klass in type(exc).__mro__:
        if klass in HANDLERS:
            return HANDLERS[klass]
    return 500


class DomainError(Exception):
    """领域层错误基类(带机器可读的 code)。"""

    def __init__(self, code: str, message: str) -> None:
        super().__init__(message)
        self.code = code


class NotFoundError(DomainError):
    """目标资源不存在。"""


class PermissionDeniedError(DomainError):
    """当前调用者无权执行该操作。"""


handles(DomainError, 400)
handles(NotFoundError, 404)
handles(PermissionDeniedError, 403)


def call_route(handler: Any, payload: dict[str, Any]) -> dict[str, Any]:
    """路由调用点: 不写 try/except, 异常交给上面注册的处理器。"""
    result = handler(payload)
    return {"status": 200, "body": result}
