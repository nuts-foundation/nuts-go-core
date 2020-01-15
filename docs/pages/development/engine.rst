.. _nuts-engine-development:

Nuts engine development
***********************

Separated pieces of logic are added to the nuts service executable as *engines*. Most engines have their own git repo as well.
An engine is an instance following struct:

.. literalinclude:: ../../../engine.go
    :language: go
    :start-after: START_DOC_ENGINE_1
    :end-before: END_DOC_ENGINE_1

Once defined it can be registered:

.. code-block:: go

    RegisterEngine(&engine)

Engine monitoring
=================

Each Engine must have a diagnostic function called `Diagnostics()` which returns information on how well the engine is performing.
The information can be used for monitoring and/or debugging.
The **status** engine exposes this information at the `/status/diagnostics` endpoint in plain text format.

.. code-block:: text

    memory usage: 256m
    established connections: 20
    loaded engines: status, logging
    ...

Standalone
==========

It is possible to run a module without adding it to the main Nuts executable by defining a Go main function:

.. code-block:: go

    // engine instance
    var e = NewMyEngine()

    // the rootCmd
    var rootCmd = e.Cmd

    // a new global nuts config
    c := cfg.NewNutsGlobalConfig()

    // ignore any config prefixes for this Cmd since it is running standalone
    c.IgnoredPrefixes = append(c.IgnoredPrefixes, e.ConfigKey)

    // register all commandLine options added by this engine
    c.RegisterFlags(e)

    // load all config from parameters into global config
    if err := c.Load(); err != nil {
        panic(err)
    }

    // inject parameters from global config into config struct of engine
    if err := c.InjectIntoEngine(e); err != nil {
        panic(err)
    }

    // check configuration on engine
    if err := e.Configure(); err != nil {
        panic(err)
    }

    // execute comand
    rootCmd.Execute()
