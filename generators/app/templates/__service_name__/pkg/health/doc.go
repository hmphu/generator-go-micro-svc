/*
Package health provides a simple health monitoring server.

To create a health service simply create a new health.Service using New,
register any probes using the RegisterProbe function and finally calling
ListenAndServe in a separate go routine:

    go func() {
      healthService := health.New()

      probe := createProbe()
      healthService.RegisterProbe("sample", probe)

      // This will block unless something failed
      err := healthService.ListenAndServe(":7000")
    }()

To implement a Probe simply adhere to the Probe interface. Use type assertion
when using a different interface:

    store := createStore() // <- store contains a interface which does not implement health.Probe
    if probe, ok := interface{}(store).(health.Probe); ok {
      healthService.RegisterProbe("store", probe)
    }

For example.

MongDB session probe, including trying to recover from an io.EOF error:

    type MongoStore struct {
      session *mgo.Session
    }

    func (s *MongoStore) Healthy() error {
      err := s.session.Ping()
      if err != nil {
        if err == io.EOF {
          s.session.Refresh()
        }

        return err
      }

      return nil
    }
*/
package health
