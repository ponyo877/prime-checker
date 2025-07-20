import React from 'react'
import { usePrimeChecksList } from '../generated-client/primeApi'
import PrimeResultItem from './PrimeResultItem'

const PrimeResultsList: React.FC = () => {
  const { data, isLoading, error } = usePrimeChecksList()

  if (isLoading) {
    return <div>Loading results...</div>
  }

  if (error) {
    return <div style={{ color: 'red' }}>Error loading results: {error.message}</div>
  }

  if (!data?.data.items || data.data.items.length === 0) {
    return <div>No prime check results yet. Try checking a number!</div>
  }

  return (
    <div>
      <h2>Results</h2>
      <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
        {data.data.items.map((item: any) => (
          <PrimeResultItem key={item.id} primeCheck={item} />
        ))}
      </div>
    </div>
  )
}

export default PrimeResultsList