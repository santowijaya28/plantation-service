# Plantation Service

This is the backend service for the Plantation management system. It manages staff, gardens (kebun), care materials (bahan perawatan), commodities (komoditas), cultivations (budidaya), harvest results (hasil panen), and plant care activities (perawatan tanaman).

## Prerequisites

Before running the application, ensure you have the following installed:

*   [Go](https://go.dev/dl/) (version 1.21.1 or later)
*   [PostgreSQL](https://www.postgresql.org/download/)

## Configuration

The application configuration is located at `internal/files/etc/conf/config.ini`.

Default configuration:
```ini
[DatabaseConfig]
    Host        = "localhost"
    Port        = "5432"
    DbName      = "plantation_db"
    DbUserName  = "app"
    DbPwd       = "app"
```

Please make sure your PostgreSQL instance is running and a database named `plantation_db` exists. You may need to adjust the credentials in `config.ini` to match your local environment.

## Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/plantation-service.git
    cd plantation-service
    ```

2.  Install dependencies:
    ```bash
    go mod vendor
    ```

## Running the Application

To run the service, use the provided `Makefile`:

```bash
make run
```

This command will:
1.  Run `go mod vendor`
2.  Build the binary to `internal/bin/http`
3.  Run the binary

The server will start on port `8000`.

## API Endpoints

The service exposes the following RESTful endpoints:

### General
*   `GET /ping`: Health check

### Staff
*   `POST /staff`: Create a new staff member
*   `GET /staff`: Get a staff member by ID (query param `id`)
*   `GET /staff/all`: Get all staff members
*   `PATCH /staff`: Update a staff member
*   `DELETE /staff`: Delete a staff member

### Kebun (Garden)
*   `POST /kebun`: Create a new garden
*   `GET /kebun`: Get a garden by ID
*   `GET /kebun/all`: Get all gardens

### Bahan Perawatan (Care Materials)
*   `POST /bahan-perawatan`: Add new care material
*   `GET /bahan-perawatan`: Get care material by ID
*   `GET /bahan-perawatan/all`: Get all care materials
*   `PATCH /bahan-perawatan`: Update care material

### Komoditas (Commodities)
*   `POST /komoditas`: Add new commodity
*   `GET /komoditas`: Get commodity by ID
*   `GET /komoditas/all`: Get all commodities

### Relationships & Activities
*   **Kebun - Bahan Perawatan**: Manage care materials assigned to gardens.
    *   `POST /kebun/bahan-perawatan`
    *   `GET /kebun/bahan-perawatan`
    *   `PATCH /kebun/bahan-perawatan`
    *   `GET /kebun/bahan-perawatan/all`
*   **Kebun - Budidaya**: Manage cultivation in gardens.
    *   `POST /kebun/budidaya`
    *   `GET /kebun/budidaya`
    *   `PATCH /kebun/budidaya`
    *   `GET /kebun/budidaya/all`

### Hasil Panen (Harvest)
*   `POST /hasil-panen`: Record harvest result
*   `GET /hasil-panen`: Get harvest result by ID
*   `GET /hasil-panen/all`: Get all harvest results
*   `PATCH /hasil-panen`: Update harvest result

### Perawatan Tanaman (Plant Care)
*   `POST /perawatan`: Record plant care activity
*   `GET /perawatan`: Get plant care activity by ID
*   `PATCH /perawatan`: Update plant care activity
*   `GET /perawatan/all`: Get all plant care activities

## Project Structure

*   `cmd/`: Entry point (though currently `internal/cmd/http/app.go` is used).
*   `internal/`:
    *   `bin/`: Compiled binaries.
    *   `cmd/`: Main application logic.
    *   `src/`: Source code including:
        *   `config/`: Configuration loading.
        *   `delivery/`: HTTP handlers.
        *   `repository/`: Database interactions.
        *   `usecase/`: Business logic.
    *   `files/`: Configuration files.
*   `pkg/`: Shared packages.
