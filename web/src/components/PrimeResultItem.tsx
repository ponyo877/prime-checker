import React from 'react'
import {
  CheckCircleIcon,
  XCircleIcon,
  QuestionMarkCircleIcon,
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
  href?: string
  icon: React.ElementType
  label: string
  disabled?: boolean
}

const ActionButton: React.FC<ActionButtonProps> = ({ href, icon: Icon, label, disabled = false }) => {
  const baseClasses = "inline-flex items-center rounded-md px-3 py-2 text-sm font-semibold shadow-sm ring-1 ring-inset"
  const enabledClasses = "bg-white text-gray-900 ring-gray-300 hover:bg-gray-50"
  const disabledClasses = "bg-gray-50 text-gray-400 ring-gray-200 cursor-not-allowed"

  const buttonClasses = `${baseClasses} ${disabled ? disabledClasses : enabledClasses}`
  const iconClasses = `mr-1.5 -ml-0.5 size-4 ${disabled ? 'text-gray-300' : 'text-gray-400'}`

  if (disabled || !href) {
    return (
      <span className="ml-3">
        <button
          disabled
          className={buttonClasses}
        >
          <Icon className={iconClasses} />
          {label}
        </button>
      </span>
    )
  }

  return (
    <span className="ml-3">
      <a
        href={href}
        target="_blank"
        rel="noopener noreferrer"
        className={buttonClasses}
      >
        <Icon className={iconClasses} />
        {label}
      </a>
    </span>
  )
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
      return { text: 'Unknown', icon: QuestionMarkCircleIcon, iconColor: 'text-gray-400' }
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
            <h2 className="text-2xl/7 font-bold text-gray-900 sm:text-3xl sm:tracking-tight truncate min-w-0">
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
          <ActionButton
            href={primeCheck.message_id ? `http://localhost:8025/search?q=message-id:${primeCheck.message_id}` : undefined}
            icon={EnvelopeIcon}
            label="Email"
            disabled={!primeCheck.message_id}
          />

          <ActionButton
            href={primeCheck.trace_id ? `http://localhost:16686/trace/${primeCheck.trace_id}` : undefined}
            icon={EyeIcon}
            label="Trace"
            disabled={!primeCheck.trace_id}
          />
        </div>
      </div>
    </div>
  )
}

export default PrimeResultItem