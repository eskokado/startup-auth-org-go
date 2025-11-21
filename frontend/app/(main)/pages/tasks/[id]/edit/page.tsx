"use client"
import TaskForm from '../../TaskForm'
import { useEffect, useState } from 'react'
import { orgApi, Organization } from '../../../../../services/org'
import { taskApi, Task } from '../../../../../services/tasks'
import { useRouter } from 'next/navigation'

export default function EditTaskPage({ params }: { params: { id: string } }) {
  const [org, setOrg] = useState<Organization | null>(null)
  const [task, setTask] = useState<Task | null>(null)
  const router = useRouter()
  const userId = typeof window !== 'undefined' ? localStorage.getItem('user-id') || '' : ''

  useEffect(() => {
    const load = async () => {
      if (!userId) return
      const o = await orgApi.getPersonal(userId)
      setOrg(o)
      const list = await taskApi.list(o.id)
      const found = list.find(t => t.id === params.id) || null
      setTask(found)
    }
    load()
  }, [userId, params.id])

  if (!org || !task) return null

  return <TaskForm title={`Editar: ${task.title}`} initial={{ title: task.title, description: task.description, status: task.status }} onSubmit={async (v) => { await taskApi.updateStatus(task.id, org.id, v.title, v.description, v.status); router.push('/pages/tasks') }} />
}