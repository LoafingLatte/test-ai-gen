import { useState } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <div className="max-w-2xl mx-auto px-4 py-16">
        <div className="text-center">
          <h1 className="text-5xl font-bold text-gray-900 mb-4">
            Vite + React + Tailwind
          </h1>
          <p className="text-xl text-gray-600 mb-8">
            Your React + Tailwind project is ready to go! 🚀
          </p>
        </div>

        <div className="bg-white rounded-lg shadow-lg p-8 mb-8">
          <div className="flex items-center justify-center mb-6">
            <div className="text-6xl">⚛️</div>
          </div>
          
          <p className="text-center text-gray-700 mb-8">
            Click the button below to test React state with Tailwind styling
          </p>
          
          <div className="flex justify-center">
            <button
              onClick={() => setCount((count) => count + 1)}
              className="px-8 py-3 bg-indigo-600 text-white font-semibold rounded-lg hover:bg-indigo-700 transition-colors duration-200 shadow-md hover:shadow-lg"
            >
              Count: {count}
            </button>
          </div>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">⚡ Vite</h3>
            <p className="text-gray-600">Lightning-fast build tool</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">⚛️ React</h3>
            <p className="text-gray-600">Library for UI components</p>
          </div>
          <div className="bg-white rounded-lg shadow p-6">
            <h3 className="text-lg font-semibold text-gray-900 mb-2">🎨 Tailwind</h3>
            <p className="text-gray-600">Utility-first CSS framework</p>
          </div>
        </div>

        <div className="mt-12 p-6 bg-blue-50 border-l-4 border-blue-600 rounded">
          <p className="text-blue-900">
            <span className="font-bold">Next steps:</span> Edit <code className="bg-white px-2 py-1 rounded">src/App.jsx</code> and save to test HMR
          </p>
        </div>
      </div>
    </div>
  )
}

export default App
