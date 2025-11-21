import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

const isTokenValid = (token: string): boolean => {
    try {
        const parts = token.split('.');
        if (parts.length !== 3) return false;
        const payload = JSON.parse(atob(parts[1]));
        const exp = typeof payload.exp === 'number' ? payload.exp : 0;
        const now = Math.floor(Date.now() / 1000);
        return exp > now;
    } catch {
        return false;
    }
};

export const useAuth = (redirectPath = '/auth/login') => {
    const router = useRouter();
    const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);

    useEffect(() => {
        const checkAuth = () => {
            const token = localStorage.getItem('access-token');
            const authStatus = !!token && isTokenValid(token!);
            setIsAuthenticated(authStatus);
            if (!authStatus && redirectPath) {
                router.push(redirectPath);
            }
        };

        // Verificação no client-side
        if (typeof window !== 'undefined') {
            checkAuth();
        }
    }, [router, redirectPath]);

    return isAuthenticated;
};
