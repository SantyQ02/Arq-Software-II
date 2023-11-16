import { useEffect } from "react"
import { TrashIcon, RefreshIcon, StopIcon, PlayIcon } from '@heroicons/react/solid'
import Image from "next/image"

import { startContainer, restartContainer, deleteContainer, stopContainer } from "@/lib/api/containers"


export default function ContainersList({ containers, refresh, setRefresh }) {

  const handleStartContainer = async (id) => {
    await startContainer(id)
    setRefresh(!refresh)
  }

  const handleDeleteContainer = async (id) => {
    await deleteContainer(id)
    setRefresh(!refresh)
  }

  const handleReStartContainer = async (id) => {
    await restartContainer(id)
    setRefresh(!refresh)
  }

  const handleStopContainer = async (id) => {
    await stopContainer(id)
    setRefresh(!refresh)
  }


  return (
    <div className="divide-y-2 divide-gray-200">
      <div className="grid grid-cols-12 gap-x-6 py-2 font-bold text-sm">
        <div className="col-span-3 flex gap-x-4 items-center">

          <div className="min-w-0 flex-auto font-bold">
            <p className="text-sm font-bold leading-6 text-gray-900">Container Name</p>
            <div className="flex">

              <p className="mt-1 truncate text-xs leading-5 text-gray-500">Container ID</p>
              {/* <p className="ml-1 mt-1 truncate text-xs leading-5 text-gray-500"> ~ {container.role.toUpperCase()}</p> */}
            </div>
          </div>
        </div>
        <div className=" text-end">
          CPU %
        </div>
        <div className="col-span-2 text-end">
          Memory Usage / Limit
        </div>
        <div className="col-span-2 text-end">
          Memory %
        </div>
        <div className="col-span-2 text-end">
          Net I / O
        </div>

        <div className="col-span-2 flex justify-end gap-3">
          Actions
        </div>
      </div>
      <ul role="list" className="divide-y divide-gray-100 pt-3">
        {containers != null && containers.map((container) => (
          <li key={container.container_id} className="grid grid-cols-12 gap-x-6 py-2 text-sm">
            <div className="col-span-3 flex gap-x-4 items-center">
              <Image
                src={'/images/docker_container_icon.png'}
                width={35}
                height={35}
              />
              <div className="min-w-0 flex-auto">
                <p className="text-sm font-semibold leading-6 text-gray-900">{container.name.toLowerCase()} {container.memory_usage == '0B' ? "(Paused)" : ""}</p>
                <div className="flex">

                  <p className="mt-1 truncate text-xs leading-5 text-gray-500">{container.container_id}</p>
                  {/* <p className="ml-1 mt-1 truncate text-xs leading-5 text-gray-500"> ~ {container.role.toUpperCase()}</p> */}
                </div>
              </div>
            </div>
            <div className=" text-end">
              {container.cpu}
            </div>
            <div className="col-span-2 text-end">
              {container.memory_usage} / {container.memory_limit}
            </div>
            <div className="col-span-2 text-end">
              {container.memory}
            </div>
            <div className="col-span-2 text-end">
              {container.net_i} / {container.net_o}
            </div>

            <div className="col-span-2 flex justify-end gap-3">
              <PlayIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleStartContainer(container.container_id)}/>
              <StopIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true"onClick={() => handleStopContainer(container.container_id)} />
              <RefreshIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleReStartContainer(container.container_id)}/>
              <TrashIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleDeleteContainer(container.container_id)}/>
            </div>
            {/* <div className=" sm:flex sm:flex-col sm:items-center ">
              <p className="text-sm leading-6 text-gray-900">Active</p>


              {container.active ?

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
            </div> */}
          </li>
        ))}

      </ul>
    </div>

  )
  // return(<div>Hola</div>)
}
