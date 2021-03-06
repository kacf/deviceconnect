# coding: utf-8

# flake8: noqa

"""
    Device Connect

    Device facing API for managing persistent device connections.   # noqa: E501

    The version of the OpenAPI document: 1
    Contact: support@mender.io
    Generated by: https://openapi-generator.tech
"""


from __future__ import absolute_import

__version__ = "1.0.0"

# import apis into sdk package
from devices_api.api.devices_api_client import DevicesAPIClient

# import ApiClient
from devices_api.api_client import ApiClient
from devices_api.configuration import Configuration
from devices_api.exceptions import OpenApiException
from devices_api.exceptions import ApiTypeError
from devices_api.exceptions import ApiValueError
from devices_api.exceptions import ApiKeyError
from devices_api.exceptions import ApiException
# import models into sdk package
from devices_api.models.error import Error

