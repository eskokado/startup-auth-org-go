"use client"
import { useEffect, useState } from 'react'
import { Card } from 'primereact/card'
import { InputText } from 'primereact/inputtext'
import { Button } from 'primereact/button'
import api from '@/app/api/api'
import { authApi } from '@/app/services/auth'

export default function ProfilePage() {
  const [name, setName] = useState('')
  const [password, setPassword] = useState('')
  const [loading, setLoading] = useState(false)
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    if (typeof window !== 'undefined') {
      const currentName = localStorage.getItem('user-name') || ''
      setName(currentName)
    }
  }, [])

  const updateName = async () => {
    if (!userId || !name) return
    setLoading(true)
    try {
      await authApi.updateName({ user_id: userId, name })
      localStorage.setItem('user-name', name)
    } finally {
      setLoading(false)
    }
  }

  const updatePassword = async () => {
    if (!userId || !password) return
    setLoading(true)
    try {
      await api.put(`/user/password/${userId}`, { password })
      setPassword('')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="grid p-4">
      <div className="col-12 md:col-6">
        <Card title="Atualizar nome">
          <div className="flex gap-2">
            <InputText className="w-full" value={name} onChange={(e) => setName(e.target.value)} />
            <Button label="Salvar" onClick={updateName} loading={loading} />
          </div>
        </Card>
      </div>
      <div className="col-12 md:col-6">
        <Card title="Atualizar senha">
          <div className="flex gap-2">
            <InputText type="password" className="w-full" value={password} onChange={(e) => setPassword(e.target.value)} />
            <Button label="Alterar" onClick={updatePassword} loading={loading} />
          </div>
        </Card>
      </div>
    </div>
  )
}