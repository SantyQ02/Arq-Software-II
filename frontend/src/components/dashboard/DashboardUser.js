/* This example requires Tailwind CSS v2.0+ */
import { PaperClipIcon } from '@heroicons/react/solid'

export default function DashboardUserProfile({ userdata }) {
  return (
    <>{userdata !== null &&
      <>
        {/* <div>
        <h3 className="text-lg leading-6 font-medium text-gray-900">Applicant Information</h3>
        <p className="mt-1 max-w-2xl text-sm text-gray-500">Personal details and application.</p>
      </div> */}
        <div className="mt-5 border-t border-gray-200">
          <dl className="divide-y divide-gray-200">
            <div className="py-4 sm:grid sm:py-5 sm:grid-cols-3 sm:gap-4">
              <dt className="text-sm font-medium text-gray-500">User ID</dt>
              <dd className="mt-1 flex text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <span className="flex-grow">{userdata.user_id}</span>
              </dd>
            </div>
            <div className="py-4 sm:py-5 sm:grid sm:grid-cols-3 sm:gap-4">
              <dt className="text-sm font-medium text-gray-500">Full name</dt>
              <dd className="mt-1 flex text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <span className="flex-grow">{userdata.first_name.toUpperCase()} {userdata.last_name.toUpperCase()}</span>
              </dd>
            </div>
            <div className="py-4 sm:grid sm:py-5 sm:grid-cols-3 sm:gap-4">
              <dt className="text-sm font-medium text-gray-500">Role</dt>
              <dd className="mt-1 flex text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <span className="flex-grow">{userdata.role.toUpperCase()}</span>
              </dd>
            </div>
            <div className="py-4 sm:grid sm:py-5 sm:grid-cols-3 sm:gap-4">
              <dt className="text-sm font-medium text-gray-500">Email address</dt>
              <dd className="mt-1 flex text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <span className="flex-grow">{userdata.email}</span>
              </dd>
            </div>

            <div className="py-4 sm:grid sm:py-5 sm:grid-cols-3 sm:gap-4">
              <dt className="text-sm font-medium text-gray-500">Username</dt>
              <dd className="mt-1 flex text-sm text-gray-900 sm:mt-0 sm:col-span-2">
                <span className="flex-grow">{userdata.user_name}</span>
              </dd>
            </div>
          </dl>
        </div>
      </>
    }

    </>
  )
}
