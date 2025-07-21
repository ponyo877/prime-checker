import React from 'react'
import { usePrimeChecksList } from '../generated-client/primeApi'
import PrimeResultItem from './PrimeResultItem'

const PrimeResultsList: React.FC = () => {
  const { data, isLoading, error } = usePrimeChecksList()

  if (isLoading) {
    return <div className="text-center text-gray-600 py-12">Loading results...</div>
  }

  if (error) {
    return <div className="text-red-600 text-sm">Error loading results: {error.message}</div>
  }

  if (!data?.data.items || data.data.items.length === 0) {
    return <div className="text-center text-gray-600 italic py-12 px-6">No results yet. Enter a number above to get started.</div>
  }

  return (
    <div className="mt-12">
      <h2 className="text-2xl font-semibold text-gray-900 mb-6">Results</h2>
      <div>
        {data.data.items.map((item: any) => (
          <PrimeResultItem key={item.id} primeCheck={item} />
        ))}
      </div>
    </div>
  )
}

export default PrimeResultsList