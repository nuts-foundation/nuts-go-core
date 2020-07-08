.. _nuts-metrics-development:

Adding metrics
**************

Please follow the manual at https://prometheus.io/docs/guides/go-application/

For now we'll use ``promauto`` to register metrics to the prometheus registry. As convention all custom metrics should start with ``nuts_``. If metrics could be interpreted as it came from multiple engines, add the engine name as prefix as well, eg: ``nuts_crypto_``.
