"use client"
import { ReactNode } from 'react'
import { useAuth } from '@/hooks/useAuth'

export default function AuthGuard({ children }: { children: ReactNode }) {
  const isAuth = useAuth('/auth/login')
  if (isAuth === false) return null
  return <>{children}</>
}