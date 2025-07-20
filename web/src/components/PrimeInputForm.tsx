import React, { useState } from 'react'
import { usePrimeChecksCreate } from '../generated-client/primeApi'

const PrimeInputForm: React.FC = () => {
  const [number, setNumber] = useState('')
  const createPrimeCheck = usePrimeChecksCreate()

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (number.trim()) {
      createPrimeCheck.mutate({ data: { number: number.trim() } })
      setNumber('')
    }
  }

  return (
    <form onSubmit={handleSubmit} style={{ marginBottom: '2rem' }}>
      <div style={{ display: 'flex', gap: '1rem', alignItems: 'center' }}>
        <input
          type="text"
          value={number}
          onChange={(e) => setNumber(e.target.value)}
          placeholder="Enter a number"
          style={{
            padding: '0.5rem',
            fontSize: '1rem',
            border: '1px solid #ccc',
            borderRadius: '4px',
            minWidth: '200px'
          }}
        />
        <button
          type="submit"
          disabled={createPrimeCheck.isPending}
          style={{
            padding: '0.5rem 1rem',
            fontSize: '1rem',
            backgroundColor: '#008CBA',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer'
          }}
        >
          {createPrimeCheck.isPending ? '...' : 'Prime?'}
        </button>
      </div>
      {createPrimeCheck.error && (
        <p style={{ color: 'red', marginTop: '0.5rem' }}>
          Error: {createPrimeCheck.error.message}
        </p>
      )}
    </form>
  )
}

export default PrimeInputForm