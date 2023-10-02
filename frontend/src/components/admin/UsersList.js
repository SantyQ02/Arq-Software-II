import { useEffect } from "react"
import { CheckCircleIcon, XCircleIcon } from '@heroicons/react/solid'
import { getUsers, updateUser } from "@/lib/api/user"


export default function UsersList({ users, setUsers }) {

  const get_users = async () => {
    const data = await getUsers()
    setUsers(data)
  }

  useEffect(() => {

    get_users()
  }, [])

  const handleCancelUser = async (e, user) => {
    e.preventDefault()
    await updateUser(user.user_id, user.first_name, user.last_name, user.user_name, user.role, false)
    get_users()
  }

  const handleRegisterUser = async (e, user) => {
    e.preventDefault()
    await updateUser(user.user_id, user.first_name, user.last_name, user.user_name, user.role, true)
    get_users()
  }

  return (
    <div className="">
      <ul role="list" className="divide-y divide-gray-100">
        {users != null && users.map((user) => (
          <li key={user.email} className="flex justify-between gap-x-6 py-5">
            <div className="flex gap-x-4 items-center">
              <span className="inline-block h-10 w-10 rounded-full overflow-hidden bg-gray-100">
                <svg className="h-full w-full text-gray-300" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M24 20.993V24H0v-2.996A14.977 14.977 0 0112.004 15c4.904 0 9.26 2.354 11.996 5.993zM16.002 8.999a4 4 0 11-8 0 4 4 0 018 0z" />
                </svg>
              </span>
              <div className="min-w-0 flex-auto">
                <p className="text-sm font-semibold leading-6 text-gray-900">{user.first_name.toUpperCase()} {user.last_name.toUpperCase()}</p>
                <div className="flex">

                  <p className="mt-1 truncate text-xs leading-5 text-gray-500">{user.email}</p>
                  <p className="ml-1 mt-1 truncate text-xs leading-5 text-gray-500"> ~ {user.role.toUpperCase()}</p>
                </div>
              </div>
            </div>
            <div className=" sm:flex sm:flex-col sm:items-center ">
              <p className="text-sm leading-6 text-gray-900">Active</p>


              {user.active ?

                <label class="relative inline-flex items-center cursor-pointer"
                  onClick={e => handleCancelUser(e, user)}>
                  <input type="checkbox" checked={true} class="sr-only peer" />
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                  <span class=" text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                </label>
                :
                <label class="relative inline-flex items-center cursor-pointer"
                  onClick={e => handleRegisterUser(e, user)}>
                  <input type="checkbox" checked={false} class="sr-only peer" />
                  <div class="w-11 h-6 bg-gray-200 peer-focus:outline-none peer-focus:ring-4 peer-focus:ring-indigo-300 dark:peer-focus:ring-indigo-800 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-indigo-600"></div>
                  <span class=" text-sm font-medium text-gray-900 dark:text-gray-300"></span>
                </label>
              }
            </div>
          </li>
        ))}

      </ul>
    </div>

  )
  // return(<div>Hola</div>)
}
