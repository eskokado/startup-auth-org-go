import api from "../api/api"

export type Task = {
  id: string
  organization_id: string
  title: string
  description: string
  status: string
}

export const taskApi = {
  list: async (organizationId: string): Promise<Task[]> => {
    const res = await api.get(`/tasks`, { params: { organization_id: organizationId } })
    return res.data as Task[]
  },
  create: async (organizationId: string, title: string, description: string) => {
    const res = await api.post('/tasks', { organization_id: organizationId, title, description })
    return res.data as { id: string }
  },
  updateStatus: async (taskId: string, organizationId: string, title: string, description: string, status: string) => {
    const res = await api.put('/tasks/status', { task_id: taskId, organization_id: organizationId, title, description, status })
    return res.data as { status: string }
  },
  delete: async (id: string) => {
    const res = await api.delete(`/tasks/${id}`)
    return res.data as { deleted: boolean }
  }
}