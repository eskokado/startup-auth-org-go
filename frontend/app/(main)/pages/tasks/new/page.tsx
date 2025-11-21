"use client"
import TaskForm from '../TaskForm'
import { useEffect, useState } from 'react'
import { orgApi, Organization } from '../../../../services/org'
import { taskApi } from '../../../../services/tasks'
import { useRouter } from 'next/navigation'

export default function NewTaskPage() {
  const [org, setOrg] = useState<Organization | null>(null)
  const router = useRouter()
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    const load = async () => {
      if (!userId) return
      const data = await orgApi.getPersonal(userId)
      setOrg(data)
    }
    load()
  }, [userId])

  if (!org) return null

  return <TaskForm title="Nova tarefa" onSubmit={async (v) => { await taskApi.create(org.id, v.title, v.description); router.push('/pages/tasks') }} />
}