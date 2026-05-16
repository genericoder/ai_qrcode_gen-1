import { useState } from 'react'
import './App.css'

const API_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:9999'

function App() {
  const [content, setContent] = useState('')
  const [qrImage, setQrImage] = useState('')
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  const generateQR = async () => {
    setError('')
    setSuccess('')
    setQrImage('')

    if (!content.trim()) {
      setError('Please enter some content')
      return
    }

    try {
      const response = await fetch(`${API_URL}/api/generate`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content }),
      })
      const data = await response.json()

      if (data.success && data.data && data.data.image) {
        setQrImage(data.data.image)
        setSuccess('QR Code generated successfully!')
      } else {
        setError(data.message || 'Failed to generate QR code')
      }
    } catch (err) {
      setError('Failed to connect to server. Make sure the server is running.')
    }
  }

  const handleKeyPress = (e) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      generateQR()
    }
  }

  return (
    <div className="container">
      <h1>QR Code Generator</h1>
      <div className="input-group">
        <label htmlFor="content">Enter content (URL, text, etc.)</label>
        <textarea
          id="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          onKeyDown={handleKeyPress}
          placeholder="https://example.com or Hello World"
        />
      </div>
      <button onClick={generateQR}>Generate QR Code</button>
      <div id="result">
        {qrImage && (
          <div id="qrcode">
            <img src={`data:image/png;base64,${qrImage}`} alt="QR Code" />
          </div>
        )}
        {error && <p className="error">{error}</p>}
        {success && <p className="success-message">{success}</p>}
      </div>
    </div>
  )
}

export default App
