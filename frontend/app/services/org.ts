import api from "../api/api"

export type Organization = { id: string; name: string }

export const orgApi = {
  getPersonal: async (ownerId: string): Promise<Organization> => {
    const res = await api.get(`/org/personal/${ownerId}`)
    return res.data as Organization
  },
  invite: async (organizationId: string, inviterId: string, email: string) => {
    const res = await api.post('/org/invite', { organization_id: organizationId, inviter_id: inviterId, email })
    return res.data as { token: string }
  },
  accept: async (token: string, userId: string) => {
    const res = await api.post('/org/invite/accept', { token, user_id: userId })
    return res.data as { status: string }
  }
}