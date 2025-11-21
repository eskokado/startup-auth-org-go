/* eslint-disable @next/next/no-img-element */

import Link from 'next/link';
import { classNames } from 'primereact/utils';
import React, { forwardRef, useContext, useImperativeHandle, useRef, useEffect, useState } from 'react';
import { AppTopbarRef } from '@/types';
import { LayoutContext } from './context/layoutcontext';
import { authApi } from '@/app/services/auth';
import { useRouter } from 'next/navigation';

const AppTopbar = forwardRef<AppTopbarRef>((props, ref) => {
    const { layoutConfig, layoutState, onMenuToggle, showProfileSidebar } = useContext(LayoutContext);
    const menubuttonRef = useRef(null);
    const topbarmenuRef = useRef(null);
    const topbarmenubuttonRef = useRef(null);
    const [logged, setLogged] = useState(false);
    const router = useRouter();

    useImperativeHandle(ref, () => ({
        menubutton: menubuttonRef.current,
        topbarmenu: topbarmenuRef.current,
        topbarmenubutton: topbarmenubuttonRef.current
    }));

    useEffect(() => {
        if (typeof window !== 'undefined') {
            setLogged(!!localStorage.getItem('access-token'));
        }
    }, [layoutState.profileSidebarVisible]);

    return (
        <div className="layout-topbar">
            <Link href="/" className="layout-topbar-logo">
                <img src={`/layout/images/logo-${layoutConfig.colorScheme !== 'light' ? 'white' : 'dark'}.svg`} width="47.22px" height={'35px'} alt="logo" />
                <span>SAKAI</span>
            </Link>

            <button ref={menubuttonRef} type="button" className="p-link layout-menu-button layout-topbar-button" onClick={onMenuToggle}>
                <i className="pi pi-bars" />
            </button>

            <button ref={topbarmenubuttonRef} type="button" className="p-link layout-topbar-menu-button layout-topbar-button" onClick={showProfileSidebar}>
                <i className="pi pi-ellipsis-v" />
            </button>

            <div ref={topbarmenuRef} className={classNames('layout-topbar-menu', { 'layout-topbar-menu-mobile-active': layoutState.profileSidebarVisible })}>
                {logged && (
                    <>
                        <Link href="/pages/profile">
                            <button type="button" className="p-link layout-topbar-button">
                                <i className="pi pi-user"></i>
                                <span>Meu Perfil</span>
                            </button>
                        </Link>
                        <button
                            type="button"
                            className="p-link layout-topbar-button"
                            onClick={async () => { await authApi.logout(); router.push('/auth/login'); }}
                        >
                            <i className="pi pi-sign-out"></i>
                            <span>Logout</span>
                        </button>
                    </>
                )}
                {!logged && (
                    <Link href="/auth/login">
                        <button type="button" className="p-link layout-topbar-button">
                            <i className="pi pi-sign-in"></i>
                            <span>Login</span>
                        </button>
                    </Link>
                )}
            </div>
        </div>
    );
});

AppTopbar.displayName = 'AppTopbar';

export default AppTopbar;
