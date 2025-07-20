import React from 'react'
import { PrimeCheck } from '../generated-client/primeApi.schemas'

interface PrimeResultItemProps {
  primeCheck: PrimeCheck
}

const PrimeResultItem: React.FC<PrimeResultItemProps> = ({ primeCheck }) => {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleString()
  }

  const getStatusBadge = (status?: string) => {
    if (!status) return null

    const statusColors = {
      completed: '#28a745',
      processing: '#ffc107',
      failed: '#dc3545'
    }

    return (
      <span
        style={{
          backgroundColor: statusColors[status as keyof typeof statusColors] || '#6c757d',
          color: 'white',
          padding: '0.25rem 0.5rem',
          borderRadius: '12px',
          fontSize: '0.75rem',
          fontWeight: 'bold'
        }}
      >
        {status.toUpperCase()}
      </span>
    )
  }

  const getPrimeResult = (isPrime?: boolean) => {
    if (isPrime === undefined) return 'Pending'
    return isPrime ? '✅ Prime' : '❌ Not Prime'
  }

  return (
    <div
      style={{
        border: '1px solid #ddd',
        borderRadius: '8px',
        padding: '1rem',
        backgroundColor: '#f9f9f9'
      }}
    >
      <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: '0.5rem' }}>
        <div>
          <h3 style={{ margin: '0 0 0.5rem 0', fontSize: '1.25rem' }}>
            Number: {primeCheck.number}
          </h3>
          <p style={{ margin: '0 0 0.5rem 0', fontSize: '1.1rem', fontWeight: 'bold' }}>
            Result: {getPrimeResult(primeCheck.is_prime)}
          </p>
        </div>
        {getStatusBadge(primeCheck.status)}
      </div>

      <div style={{ fontSize: '0.9rem', color: '#666', marginBottom: '1rem' }}>
        <p style={{ margin: '0.25rem 0' }}>ID: {primeCheck.id}</p>
        <p style={{ margin: '0.25rem 0' }}>Created: {formatDate(primeCheck.created_at)}</p>
        {primeCheck.trace_id && (
          <p style={{ margin: '0.25rem 0' }}>Trace ID: {primeCheck.trace_id}</p>
        )}
        {primeCheck.message_id && (
          <p style={{ margin: '0.25rem 0' }}>Message ID: {primeCheck.message_id}</p>
        )}
      </div>

      <div style={{ display: 'flex', gap: '0.5rem' }}>
        {primeCheck.message_id && (
          <a
            href={`http://localhost:8025/search?query=${primeCheck.message_id}`}
            target="_blank"
            rel="noopener noreferrer"
            style={{
              padding: '0.5rem 1rem',
              backgroundColor: '#17a2b8',
              color: 'white',
              textDecoration: 'none',
              borderRadius: '4px',
              fontSize: '0.875rem'
            }}
          >
            📧 Email
          </a>
        )}
        {primeCheck.trace_id && (
          <a
            href={`http://localhost:16686/trace/${primeCheck.trace_id}`}
            target="_blank"
            rel="noopener noreferrer"
            style={{
              padding: '0.5rem 1rem',
              backgroundColor: '#6f42c1',
              color: 'white',
              textDecoration: 'none',
              borderRadius: '4px',
              fontSize: '0.875rem'
            }}
          >
            📊 Trace
          </a>
        )}
      </div>
    </div>
  )
}

export default PrimeResultItem