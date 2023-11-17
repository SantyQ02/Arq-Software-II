import { Oval } from 'react-loader-spinner'
import { TrashIcon, RefreshIcon, StopIcon, PlayIcon } from '@heroicons/react/solid'
import Image from "next/image"

import { startContainer, restartContainer, deleteContainer, stopContainer } from "@/lib/api/containers"
import { useState } from 'react'

const ContainerRow = ({ container }) => {
  const [isLoading, setIsLoading] = useState(false)

  const actions = {
    start: async (id) => await startContainer(id),
    delete: async (id) => await deleteContainer(id),
    stop: async (id) => await stopContainer(id),
    restart: async (id) => await restartContainer(id),

  }
  const handleContainer = async (action, id) => {
    setIsLoading(true)
    await actions[action](id)
    setIsLoading(false)
  }

  return (
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
        {
          !isLoading ?
            <>
              <PlayIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleContainer("start", container.container_id)} />
              <StopIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleContainer("stop", container.container_id)} />
              <RefreshIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleContainer("restart", container.container_id)} />
              <TrashIcon className="flex-shrink-0 h-5 w-5 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleContainer("delete", container.container_id)} />
            </>
            :
            <Oval
              type="Oval"
              color="#000"
              width={20}
              height={20}
            />
        }

      </div>
    </li>
  )
}


export default function ContainersList({ containers, refresh, setRefresh }) {






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
          <ContainerRow container={container} />
        ))
        }

      </ul>
    </div>

  )
  // return(<div>Hola</div>)
}
