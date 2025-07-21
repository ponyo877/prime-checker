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
    <div className="mb-12">
      <form onSubmit={handleSubmit} className="flex gap-3 mb-4">
        <input
          type="text"
          value={number}
          onChange={(e) => setNumber(e.target.value)}
          placeholder="Enter a number"
          className="flex-1 px-4 py-3 text-base border-2 border-gray-300 rounded-lg bg-white font-sans transition-colors focus:outline-none focus:border-blue-500 focus:ring-2 focus:ring-blue-100 placeholder-gray-400"
        />
        <button
          type="submit"
          disabled={createPrimeCheck.isPending}
          className="px-6 py-3 text-base font-medium bg-gray-900 text-white border-none rounded-lg cursor-pointer transition-colors hover:bg-gray-700 disabled:bg-gray-400 disabled:cursor-not-allowed whitespace-nowrap"
        >
          {createPrimeCheck.isPending ? 'Checking...' : 'Check'}
        </button>
      </form>
      {createPrimeCheck.error && (
        <div className="text-red-600 text-sm mt-2">
          {createPrimeCheck.error.message}
        </div>
      )}
    </div>
  )
}

export default PrimeInputForm