import './App.css'
import Payment from './components/Payment'

function App() {
  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header Section */}
      <div className="bg-gradient-to-r from-green-500 to-green-600 text-white py-12 px-4">
        <div className="max-w-4xl mx-auto text-center">
          <div className="flex justify-center gap-4 mb-6">
            <div className="text-3xl">🎯</div>
            <div className="text-3xl">⚙️</div>
          </div>
          <h1 className="text-4xl font-bold mb-3">PayFlow Demo</h1>
          <p className="text-lg opacity-90">
            Experience seamless payment processing with modern React components<br />
            and beautiful UI design
          </p>
        </div>
      </div>

      {/* Main Content */}
      <div className="max-w-5xl mx-auto px-4 py-12">
        {/* Features Section */}
        <div className="text-center mb-12">
          <h2 className="text-3xl font-bold text-gray-900 mb-3">Modern Payment Solutions</h2>
          <p className="text-gray-600 max-w-2xl mx-auto">
            Built with React and Vite for lightning-fast performance. Secure, responsive, and user-friendly payment forms that adapt to any device.
          </p>
        </div>

        {/* Feature Cards */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
          <div className="bg-white rounded-lg shadow-md p-8 text-center">
            <div className="text-4xl mb-4">🔒</div>
            <h3 className="text-xl font-semibold text-gray-900 mb-2">Secure Processing</h3>
            <p className="text-gray-600">
              End-to-end encryption ensures your payment data is always protected and secure
            </p>
          </div>
          <div className="bg-white rounded-lg shadow-md p-8 text-center">
            <div className="text-4xl mb-4">⚡</div>
            <h3 className="text-xl font-semibold text-gray-900 mb-2">Lightning Fast</h3>
            <p className="text-gray-600">
              Powered by Vite and React for instant load times and smooth interactions
            </p>
          </div>
          <div className="bg-white rounded-lg shadow-md p-8 text-center">
            <div className="text-4xl mb-4">📱</div>
            <h3 className="text-xl font-semibold text-gray-900 mb-2">Mobile Ready</h3>
            <p className="text-gray-600">
              Fully responsive design that works perfectly on all devices and screen sizes
            </p>
          </div>
        </div>

        {/* Interactive Demo Section */}
        <div className="bg-white rounded-lg shadow-md p-12 text-center mb-12">
          <h3 className="text-2xl font-semibold text-gray-900 mb-4">Interactive Demo</h3>
          <button className="bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-6 rounded-lg transition-colors">
            👉 Click Counter: 0
          </button>
          <p className="text-gray-600 text-sm mt-6">
            Edit, integrate, and easily to test Hall Module Requirement
          </p>
        </div>
      </div>

      {/* Payment Form Section */}
      <div className="bg-gray-800 text-white py-16 px-4">
        <div className="max-w-2xl mx-auto text-center mb-12">
          <h2 className="text-3xl font-bold mb-3">Try Our Payment Form</h2>
          <p className="text-gray-400">
            Experience our beautifully designed payment form with real-time validation, secure input handling, and smooth animations.
          </p>
        </div>

        <div className="max-w-md mx-auto">
          <Payment />
        </div>

        <div className="mt-12 text-center text-sm text-gray-500">
          <p>👇 This is a demo form - no real payments will be processed</p>
          <div className="flex justify-center gap-6 mt-4 flex-wrap">
            <span className="text-green-400">✓ Secure & Compliant</span>
            <span className="text-green-400">• Mobile Optimized</span>
            <span className="text-green-400">🛡️ Data Protection</span>
          </div>
        </div>
      </div>

      {/* Footer */}
      <div className="bg-gray-900 text-gray-400 py-12 px-4">
        <div className="max-w-5xl mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8 mb-8 text-center md:text-left">
            <div>
              <h4 className="text-white font-semibold mb-2">PayFlow Demo</h4>
              <p className="text-sm">A modern payment form built with React, Vite, and Tailwind CSS. Showcasing best practices of UI/UX design.</p>
            </div>
            <div>
              <h4 className="text-white font-semibold mb-2">Technologies</h4>
              <ul className="text-sm space-y-1">
                <li>▪ React 18</li>
                <li>⚡ Vite</li>
                <li>🎨 Tailwind CSS</li>
                <li>🌐 Modern JavaScript</li>
              </ul>
            </div>
            <div>
              <h4 className="text-white font-semibold mb-2">Learn More</h4>
              <ul className="text-sm space-y-1">
                <li><a href="#" className="text-green-400 hover:text-green-300">React Docs</a></li>
                <li><a href="#" className="text-green-400 hover:text-green-300">Vite Guide</a></li>
                <li><a href="#" className="text-green-400 hover:text-green-300">Tailwind Utilities</a></li>
                <li><a href="#" className="text-green-400 hover:text-green-300">Modern JavaScript</a></li>
              </ul>
            </div>
          </div>
          <div className="border-t border-gray-700 pt-8 text-center text-sm">
            <p>Built with ❤️ using modern web technologies</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
