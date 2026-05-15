import { useState } from 'react'
import { QRCodeCanvas } from 'qrcode.react'
import './App.css'

function App() {
  const [content, setContent] = useState('')
  const [qrData, setQrData] = useState('')
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  const generateQR = async () => {
    setError('')
    setSuccess('')
    setQrData('')

    if (!content.trim()) {
      setError('Please enter some content')
      return
    }

    try {
      const response = await fetch('http://localhost:9999/api/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ content }),
      })
      const data = await response.json()

      if (data.success) {
        setQrData(data.data)
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
        {qrData && (
          <div id="qrcode">
            <QRCodeCanvas value={qrData} size={300} level="H" />
          </div>
        )}
        {error && <p className="error">{error}</p>}
        {success && <p className="success-message">{success}</p>}
      </div>
    </div>
  )
}

export default App
