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

    const baseClass = "px-3 py-1.5 rounded-2xl text-xs font-medium uppercase tracking-wider"

    const statusClass = {
      completed: "bg-green-100 text-green-800",
      processing: "bg-yellow-100 text-yellow-800",
      failed: "bg-red-100 text-red-800"
    }[status] || "bg-gray-100 text-gray-800"

    return (
      <span className={`${baseClass} ${statusClass}`}>
        {status}
      </span>
    )
  }

  const getPrimeResult = (isPrime?: boolean) => {
    if (isPrime === undefined) {
      return { text: 'Checking...', className: 'text-gray-600' }
    }
    return isPrime
      ? { text: 'Prime', className: 'text-green-600' }
      : { text: 'Not Prime', className: 'text-red-600' }
  }

  const primeResult = getPrimeResult(primeCheck.is_prime)

  return (
    <div className="border border-gray-200 rounded-xl p-6 mb-4 bg-white transition-colors hover:border-gray-300">
      <div className="flex justify-between items-start mb-4">
        <div>
          <h3 className="text-3xl font-bold text-gray-900 mb-2">{primeCheck.number}</h3>
          <p className={`text-base font-medium ${primeResult.className}`}>
            {primeResult.text}
          </p>
        </div>
        {getStatusBadge(primeCheck.status)}
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 mb-5 text-sm text-gray-600">
        <div className="flex justify-between">
          <span className="font-medium">ID</span>
          <span>{primeCheck.id}</span>
        </div>
        <div className="flex justify-between">
          <span className="font-medium">Created</span>
          <span>{formatDate(primeCheck.created_at)}</span>
        </div>
        {primeCheck.trace_id && (
          <div className="flex justify-between">
            <span className="font-medium">Trace ID</span>
            <span>{primeCheck.trace_id.slice(0, 8)}...</span>
          </div>
        )}
        {primeCheck.message_id && (
          <div className="flex justify-between">
            <span className="font-medium">Message ID</span>
            <span>{primeCheck.message_id}</span>
          </div>
        )}
      </div>

      <div className="flex gap-3 flex-wrap">
        {primeCheck.message_id && (
          <a
            href={`http://localhost:8025/search?query=${primeCheck.message_id}`}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-2 px-4 py-2 text-sm font-medium text-blue-700 bg-blue-50 border border-blue-200 rounded-md hover:opacity-80 transition-opacity"
          >
            Email
          </a>
        )}
        {primeCheck.trace_id && (
          <a
            href={`http://localhost:16686/trace/${primeCheck.trace_id}`}
            target="_blank"
            rel="noopener noreferrer"
            className="inline-flex items-center gap-2 px-4 py-2 text-sm font-medium text-gray-700 bg-gray-50 border border-gray-200 rounded-md hover:opacity-80 transition-opacity"
          >
            Trace
          </a>
        )}
      </div>
    </div>
  )
}

export default PrimeResultItem