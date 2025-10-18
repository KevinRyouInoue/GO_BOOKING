# GO_BOOKING - Conference Ticket Booking App

A simple Go application for booking conference tickets with user input validation and concurrent ticket sending.

## Prerequisites

- Go 1.17 or higher (tested with Go 1.25.1)

## How to Start

### Option 1: Run directly
```powershell
go run .
```

### Option 2: Build and run the executable
```powershell
# Build the app
go build

# Run the executable
.\booking-app.exe
```

## How to Use

1. Start the application using one of the methods above
2. The app will greet you and show available tickets
3. Follow the prompts to enter:
   - First name (min 2 characters)
   - Last name (min 2 characters)
   - Email address (must contain @)
   - Number of tickets (must be available)
4. After booking, you'll receive a confirmation
5. The app simulates sending a ticket email (50-second delay using goroutine)

## Features

- ✅ User input validation
- ✅ Concurrent ticket processing with goroutines
- ✅ Tracks remaining tickets
- ✅ Stores booking data in memory

## Project Structure

```
.
├── main.go      # Main application logic and entry point
├── helper.go    # Input validation functions
├── go.mod       # Go module definition
└── README.md    # This file
```

## Notes

- The conference has 50 tickets total
- Ticket emails are simulated with a 50-second delay
- Currently configured to process one booking (for loop is commented out)
