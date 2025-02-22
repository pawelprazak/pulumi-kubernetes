# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'TokenReviewSpec',
    'TokenReviewStatus',
    'UserInfo',
]

@pulumi.output_type
class TokenReviewSpec(dict):
    """
    TokenReviewSpec is a description of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 token: Optional[str] = None):
        """
        TokenReviewSpec is a description of the token authentication request.
        :param Sequence[str] audiences: Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        :param str token: Token is the opaque bearer token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def token(self) -> Optional[str]:
        """
        Token is the opaque bearer token.
        """
        return pulumi.get(self, "token")


@pulumi.output_type
class TokenReviewStatus(dict):
    """
    TokenReviewStatus is the result of the token authentication request.
    """
    def __init__(__self__, *,
                 audiences: Optional[Sequence[str]] = None,
                 authenticated: Optional[bool] = None,
                 error: Optional[str] = None,
                 user: Optional['outputs.UserInfo'] = None):
        """
        TokenReviewStatus is the result of the token authentication request.
        :param Sequence[str] audiences: Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        :param bool authenticated: Authenticated indicates that the token was associated with a known user.
        :param str error: Error indicates that the token couldn't be checked
        :param 'UserInfoArgs' user: User is the UserInfo associated with the provided token.
        """
        if audiences is not None:
            pulumi.set(__self__, "audiences", audiences)
        if authenticated is not None:
            pulumi.set(__self__, "authenticated", authenticated)
        if error is not None:
            pulumi.set(__self__, "error", error)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def audiences(self) -> Optional[Sequence[str]]:
        """
        Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
        """
        return pulumi.get(self, "audiences")

    @property
    @pulumi.getter
    def authenticated(self) -> Optional[bool]:
        """
        Authenticated indicates that the token was associated with a known user.
        """
        return pulumi.get(self, "authenticated")

    @property
    @pulumi.getter
    def error(self) -> Optional[str]:
        """
        Error indicates that the token couldn't be checked
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def user(self) -> Optional['outputs.UserInfo']:
        """
        User is the UserInfo associated with the provided token.
        """
        return pulumi.get(self, "user")


@pulumi.output_type
class UserInfo(dict):
    """
    UserInfo holds the information about the user needed to implement the user.Info interface.
    """
    def __init__(__self__, *,
                 extra: Optional[Mapping[str, Sequence[str]]] = None,
                 groups: Optional[Sequence[str]] = None,
                 uid: Optional[str] = None,
                 username: Optional[str] = None):
        """
        UserInfo holds the information about the user needed to implement the user.Info interface.
        :param Mapping[str, Sequence[str]] extra: Any additional information provided by the authenticator.
        :param Sequence[str] groups: The names of groups this user is a part of.
        :param str uid: A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        :param str username: The name that uniquely identifies this user among all active users.
        """
        if extra is not None:
            pulumi.set(__self__, "extra", extra)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def extra(self) -> Optional[Mapping[str, Sequence[str]]]:
        """
        Any additional information provided by the authenticator.
        """
        return pulumi.get(self, "extra")

    @property
    @pulumi.getter
    def groups(self) -> Optional[Sequence[str]]:
        """
        The names of groups this user is a part of.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def uid(self) -> Optional[str]:
        """
        A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter
    def username(self) -> Optional[str]:
        """
        The name that uniquely identifies this user among all active users.
        """
        return pulumi.get(self, "username")


