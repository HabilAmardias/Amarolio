import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import RedirectPage from './pages/RedirectPage'
import './App.css'

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/:slug" element={<RedirectPage />} />
      </Routes>
    </Router>
  )
}

export default App
