import { useEffect, useState, createContext } from 'react';
import { useRouter } from 'next/router';

import { getMe } from '@/lib/api/user';
import { refresh } from '@/lib/api/auth';

const UserContext = createContext(null);

export default function LayoutContext({ title, children }) {
    const [user, setUser] = useState(null);
    const router = useRouter();

    const getUser = async () => {
        const currentUser = await getMe()
        setUser(currentUser)
        await refresh()
    }

    useEffect(() => {
        getUser()
      router.events.on('routeChangeStart', getUser);
  
      return () => {
        router.events.off('routeChangeStart', getUser);
      };
    }, []);

    return (
        <UserContext.Provider value={[user, setUser]}>
            <div>


                {children}

            </div>
        </UserContext.Provider>
    )
}

export { UserContext }