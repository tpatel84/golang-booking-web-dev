# Bookings and Reservations

This is the repository for my bookings and reservations project.

- Built in go version 1.15
- Uses the [chi router](https://github.com/go-chi/chi)
- Uses alex edwards [SCS session management](https://github.com/alexedwards/scs)
- Uses [nosurf](https://github.com/justinas/nosurf)

# Check code coverage in golang

go test -coverprofile=coverage.out && go tool cover -html=coverage.out