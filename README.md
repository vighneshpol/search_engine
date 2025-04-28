# 🚀 Search Engine Project

## Overview

This project aims to build a simple **Search Engine** using **Golang** for the backend and **React.js** for the frontend. The search engine loads structured data from multiple files and allows users to search through it based on various fields like `Message`, `Sender`, `Event`, etc.

### **Backend (Golang)**
- In-memory data loading from files.
- Search endpoint (`/search`) to query the dataset based on a search term.
- Implements a case-insensitive search across relevant fields in the dataset.

### **Frontend (React.js)**
- A simple user interface to input search queries.
- Displays the search results with key data points.
- Shows the time taken for query execution for transparency.

---

## 🔧 Technologies Used

| Technology      | Description                                          |
|-----------------|------------------------------------------------------|
| **Backend**     | Golang (net/http), JSON Parsing                      |
| **Frontend**    | React.js (Vite or Create React App)                  |
| **Data Handling**| JSON structured data processing                    |
| **Deployment**  | Docker & Docker Compose (Optional for future use)    |

---

## 📂 Folder Structure


# Folder Structure

- **search-engine/**
  - **backend/**
    - `main.go` (Main Go server code)
    - **utils/**
      - `parquet_reader.go` (File reading utility)
    - **parquet_files/** (Contains File 1, File 2, ..., File 16)
    - `go.mod` / `go.sum` (Go modules)
  - **frontend/**
    - **src/**
      - `App.jsx` (Main React App)
      - **components/**
        - `SearchBar.jsx` (Search component)
      - **services/**
        - `api.js` (API calling logic)
    - `package.json` (React project dependencies)
  - `README.md` (Project description)



## ✅ Work Done

1. **Backend (Golang)**:
   - Implemented data loading from files located under `parquet_files/`.
   - Set up a basic REST API (`/search`) to accept search queries and return results.
   - Implemented logic to search through important fields (`Message`, `Sender`, `Event`, etc.).
   - Applied case-insensitive matching for search queries.

2. **Frontend (React.js)**:
   - Created the basic React app with a search bar.
   - Connected the search bar to the backend API and displayed search results.
   - Added basic error handling and loading states for the search functionality.

3. **Testing**:
   - Successfully loaded files and data into memory.
   - Validated that the backend is receiving and handling search queries.
   - Search results show as expected for test queries in the terminal (still being refined).

---

## 🛠 Pending Work

1. **Search Query Issue**:
   - The search query is not returning expected results in the browser or through curl commands.
   - There may be an issue with data parsing from the `.parquet` files or mismatched fields that prevent correct results from being returned.

2. **Search Performance**:
   - Data processing and memory optimization for larger datasets.
   - Add pagination to results to handle large data sets effectively.

3. **Frontend Enhancements**:
   - Add more user-friendly feedback for no results found or errors.
   - Improve search result formatting (e.g., highlight the matched text).

4. **Error Handling**:
   - Improve error responses for invalid queries or unexpected issues with the backend API.

---

## 🧪 Testing

### Backend:
- The backend server is running at `http://localhost:8080`.
- Search functionality is tested using `curl` and Postman for various queries.
- **Known Issue**: Search results are returning zero matches for certain queries, potentially due to data structure or search field mismatch.

### Frontend:
- React app can be tested by running it on `http://localhost:3000`.
- User can enter a search query, and results are displayed (currently limited by search issue).

### Unit Testing:
- Unit tests can be written for backend functionality in Go (to be done).

---

## 📝 Notes

- **Known Issues**: The search functionality is still under investigation. While data is loading correctly, the search queries do not return the expected results in some cases.
  
- **Data Files**: The data being used for this project are simulated and stored as files (`File 1`, `File 2`, ... `File 16`). These files are expected to contain structured data for search queries.

---

## 🚀 How to Run the Project

### 1. **Backend (Golang)**:
   - Install Golang: [Golang Installation Guide](https://golang.org/doc/install)
   - Navigate to the `backend` directory and run the Go server:
     ```bash
     cd backend
     go run main.go
     ```
   - The server will be running at `http://localhost:8080`.

### 2. **Frontend (React.js)**:
   - Install Node.js and npm: [Node.js Installation](https://nodejs.org/)
   - Navigate to the `frontend` directory and install the required dependencies:
     ```bash
     cd frontend
     npm install
     ```
   - Start the React app:
     ```bash
     npm start
     ```
   - The app will be running at `http://localhost:3000`.

---

## 💬 Feedback

If you have suggestions or feedback regarding the project or encountered any issues, feel free to reach out!  
We appreciate any input on improving the search query functionality and optimizing the system for larger datasets.

