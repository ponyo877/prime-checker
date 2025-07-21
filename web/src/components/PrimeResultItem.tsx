import React from 'react'
import {
  CheckCircleIcon,
  XCircleIcon,
  ClockIcon,
  IdentificationIcon,
  CalendarIcon,
  EnvelopeIcon,
  EyeIcon
} from '@heroicons/react/20/solid'
import { PrimeCheck } from '../generated-client/primeApi.schemas'

interface PrimeResultItemProps {
  primeCheck: PrimeCheck
}

interface ActionButtonProps {
  href: string
  icon: React.ElementType
  label: string
}

const ActionButton: React.FC<ActionButtonProps> = ({ href, icon: Icon, label }) => (
  <span className="ml-3">
    <a
      href={href}
      target="_blank"
      rel="noopener noreferrer"
      className="inline-flex items-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-gray-300 ring-inset hover:bg-gray-50"
    >
      <Icon className="mr-1.5 -ml-0.5 size-4 text-gray-400" />
      {label}
    </a>
  </span>
)

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
      return { text: 'Unknown', icon: ClockIcon, iconColor: 'text-gray-400' }
    }
    return isPrime
      ? { text: 'Prime', icon: CheckCircleIcon, iconColor: 'text-green-500' }
      : { text: 'Not Prime', icon: XCircleIcon, iconColor: 'text-red-500' }
  }

  const primeResult = getPrimeResult(primeCheck.is_prime)

  return (
    <div className="bg-white border border-gray-200 rounded-lg p-6 mb-6">
      <div className="lg:flex lg:items-start lg:justify-between">
        <div className="min-w-0 flex-1">
          <div className="flex items-center gap-3">
            <h2 className="text-2xl/7 font-bold text-gray-900 sm:truncate sm:text-3xl sm:tracking-tight">
              {primeCheck.number}
            </h2>
            {getStatusBadge(primeCheck.status)}
          </div>
          <div className="mt-1 flex flex-col sm:mt-0 sm:flex-row sm:flex-wrap sm:space-x-6">
            <div className="mt-2 flex items-center text-sm text-gray-500">
              <primeResult.icon className={`mr-1.5 size-4 shrink-0 ${primeResult.iconColor}`} />
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
          </div>
        </div>
        <div className="mt-5 flex lg:mt-0 lg:ml-4">
          {primeCheck.message_id && (
            <ActionButton
              href={`http://localhost:8025/search?query=${primeCheck.message_id}`}
              icon={EnvelopeIcon}
              label="Email"
            />
          )}

          {primeCheck.trace_id && (
            <ActionButton
              href={`http://localhost:16686/trace/${primeCheck.trace_id}`}
              icon={EyeIcon}
              label="Trace"
            />
          )}
        </div>
      </div>
    </div>
  )
}

export default PrimeResultItem