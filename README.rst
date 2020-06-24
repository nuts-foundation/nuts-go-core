nuts-go-core
============

Common resources for Nuts Go modules.

.. image:: https://circleci.com/gh/nuts-foundation/nuts-go-core.svg?style=svg
    :target: https://circleci.com/gh/nuts-foundation/nuts-go-core
    :alt: Build Status

.. image:: https://codecov.io/gh/nuts-foundation/nuts-go-core/branch/master/graph/badge.svg
    :target: https://codecov.io/gh/nuts-foundation/nuts-go-core

.. image:: https://api.codeclimate.com/v1/badges/641734b46b0950436e39/maintainability
   :target: https://codeclimate.com/github/nuts-foundation/nuts-go-core/maintainability
   :alt: Maintainability

Building
------------

.. note::

    Nuts-go uses Go version >= `1.13`.

.. code-block:: shell

   go get github.com/nuts-foundation/nuts-go-core

For generating mocks
--------------------

.. code-block:: shell

   go get github.com/golang/mock/gomock
   go install github.com/golang/mock/mockgen

Then run

.. code-block:: shell

   mockgen -destination=mock/mock_oapi.go -package=mock github.com/deepmap/oapi-codegen/pkg/runtime EchoRouter
   mockgen -destination=mock/mock_echo.go -package=mock github.com/labstack/echo/v4 Context

Testing
-------

.. code-block:: shell

   go test ./...
