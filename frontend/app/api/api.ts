import axios from 'axios';

export const baseURL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

const api = axios.create({ baseURL, withCredentials: true });
interface ResponseData {
    access_token?: string;
    user?: UserData
}

interface UserData {
    id?: string;
    name?: string;
    email?: string;
}


api.interceptors.request.use((config) => {
    const token = localStorage.getItem('access-token');
    if (token) {
        config.headers!['authorization'] = `Bearer ${token}`;
    }
    return config;
});

api.interceptors.response.use(
    (response) => {
        const newAccessToken = (response.data as ResponseData)?.access_token;
        const newUserId = (response.data as ResponseData)?.user?.id
        const newUserName = (response.data as ResponseData)?.user?.name
        const newUserEmail = (response.data as ResponseData)?.user?.email

        if (newAccessToken && newUserId && newUserEmail && newUserName) {
            localStorage.setItem('access-token', newAccessToken);
            localStorage.setItem('user-id', newUserId)
            localStorage.setItem('user-name', newUserName)
            localStorage.setItem('user-email', newUserEmail)
            const payload = decodeJwt(newAccessToken);
            if (payload) {
                if (payload.org_id) localStorage.setItem('org-id', String(payload.org_id));
                if (payload.plan) localStorage.setItem('plan', String(payload.plan));
                if (payload.exp) localStorage.setItem('token-exp', String(payload.exp));
            }
        }

        return response;
    },
    (error) => {
        if (error.response) {
            if (error.response?.status === 401) {
                localStorage.removeItem('access-token');
                window.location.href = '/auth/login';
                return Promise.reject({ code: 401, messages: ['Não autorizado'] });
            }


            const errorData = error.response.data;
            let errorMessages: string[] = [];

            if (errorData.errors && Array.isArray(errorData.errors)) {
                errorMessages = errorData.errors; // Array
            } else if (errorData.error) {
                errorMessages = [errorData.error]; // Converte string em array
            } else {
                errorMessages = ['Erro desconhecido'];
            }

            const errorCode = error.response.status || 500;
            return Promise.reject({ code: errorCode, messages: errorMessages });
        }

        const errorMessage = error.response?.data?.message || 'Erro desconhecido';
        const errorCode = error.response?.status || 500;
        return Promise.reject({ code: errorCode, messages: [errorMessage] });
    }
);

export default api;
const decodeJwt = (token: string): any | null => {
    try {
        const parts = token.split('.');
        if (parts.length !== 3) return null;
        return JSON.parse(atob(parts[1]));
    } catch {
        return null;
    }
};
