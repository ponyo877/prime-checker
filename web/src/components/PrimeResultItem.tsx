import React from 'react'
import { 
  CheckCircleIcon,
  XCircleIcon,
  ClockIcon,
  IdentificationIcon,
  CalendarIcon,
  MagnifyingGlassIcon,
  EnvelopeIcon,
  EyeIcon
} from '@heroicons/react/20/solid'
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
    <div className="bg-white border border-gray-200 rounded-lg p-6 mb-6">
      <div className="lg:flex lg:items-center lg:justify-between">
        <div className="min-w-0 flex-1">
          <h2 className="text-2xl/7 font-bold text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
            {primeCheck.number}
          </h2>
          <div className="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap sm:space-x-6">
            <div className="mt-2 flex items-center text-sm text-gray-500">
              {primeResult.className === 'text-green-600' ? (
                <CheckCircleIcon className="mr-1.5 size-4 shrink-0 text-green-500" />
              ) : primeResult.className === 'text-red-600' ? (
                <XCircleIcon className="mr-1.5 size-4 shrink-0 text-red-500" />
              ) : (
                <ClockIcon className="mr-1.5 size-4 shrink-0 text-gray-400" />
              )}
              {primeResult.text}
            </div>
            <div className="mt-2 flex items-center text-sm text-gray-500">
              <IdentificationIcon className="mr-1.5 size-4 shrink-0 text-gray-400" />
              ID: {primeCheck.id}
            </div>
            <div className="mt-2 flex items-center text-sm text-gray-500">
              <CalendarIcon className="mr-1.5 size-4 shrink-0 text-gray-400" />
              {formatDate(primeCheck.created_at)}
            </div>
            {primeCheck.trace_id && (
              <div className="mt-2 flex items-center text-sm text-gray-500">
                <MagnifyingGlassIcon className="mr-1.5 size-4 shrink-0 text-gray-400" />
                Trace: {primeCheck.trace_id.slice(0, 8)}...
              </div>
            )}
          </div>
        </div>
        <div className="mt-5 flex lg:mt-0 lg:ml-4">
          {getStatusBadge(primeCheck.status)}
          
          {primeCheck.message_id && (
            <span className="ml-3">
              <a
                href={`http://localhost:8025/search?query=${primeCheck.message_id}`}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-gray-300 ring-inset hover:bg-gray-50"
              >
                <EnvelopeIcon className="mr-1.5 -ml-0.5 size-4 text-gray-400" />
                Email
              </a>
            </span>
          )}

          {primeCheck.trace_id && (
            <span className="ml-3">
              <a
                href={`http://localhost:16686/trace/${primeCheck.trace_id}`}
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                <EyeIcon className="mr-1.5 -ml-0.5 size-4" />
                View Trace
              </a>
            </span>
          )}
        </div>
      </div>
    </div>
  )
}

export default PrimeResultItem