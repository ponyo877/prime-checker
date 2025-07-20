import PrimeInputForm from './components/PrimeInputForm'
import PrimeResultsList from './components/PrimeResultsList'

function App() {
  return (
    <div style={{ padding: '2rem', maxWidth: '800px', margin: '0 auto' }}>
      <h1>Prime Checker</h1>
      <PrimeInputForm />
      <PrimeResultsList />
    </div>
  )
}

export default App