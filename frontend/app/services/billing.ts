import api from "../api/api"

export type CheckoutResponse = { checkout_url: string }

export const billingApi = {
  checkout: async (organizationId: string, plan: string, cycle: string, successUrl: string, cancelUrl: string): Promise<CheckoutResponse> => {
    const res = await api.post('/billing/checkout', { organization_id: organizationId, plan, cycle, success_url: successUrl, cancel_url: cancelUrl })
    return res.data as CheckoutResponse
  }
}