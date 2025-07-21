import PrimeInputForm from './components/PrimeInputForm'
import PrimeResultsList from './components/PrimeResultsList'

function App() {
  return (
    <div className="max-w-3xl mx-auto px-6 py-12">
      <div className="text-center mb-16">
        <h1 className="text-5xl font-bold text-gray-900 mb-4 leading-tight">
          Prime Checker
        </h1>
        <p className="text-lg text-gray-500">
          Check if numbers are prime and trace the complete processing journey through our distributed system
        </p>
      </div>
      <PrimeInputForm />
      <PrimeResultsList />
    </div>
  )
}

export default App