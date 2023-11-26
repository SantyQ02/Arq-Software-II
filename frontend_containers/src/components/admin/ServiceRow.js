import ContainersList from "@/components/admin/ContainersList"
import { createContainer } from "@/lib/api/containers"
import { useState } from "react"
import { PlusIcon, MinusIcon } from '@heroicons/react/solid'
import { Oval } from 'react-loader-spinner'

const ServiceRow = ({item, refresh, setRefresh, containers}) => {
    const [isOpen, setIsOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(false)

    const handleNewInstance = async (service) => {
        setIsLoading(true)
        await createContainer(service)
        setRefresh(!refresh)
        setIsLoading(false)
      }

    return (
        <div className="pt-5 divide-y-2 divide-dashed">
            <div className="mb-5 font-lg font-bold grid grid-cols-12 items-center">
              <div className="col-span-2">
                {item.title}

              </div>
              {
                isOpen ?
                  <MinusIcon className="flex-shrink-0 h-7 w-7 text-blue-400 cursor-pointer col-span-8" aria-hidden="true" onClick={() => setIsOpen(!isOpen)} />
                  :
                  <PlusIcon className="flex-shrink-0 h-7 w-7 text-blue-400 cursor-pointer col-span-8" aria-hidden="true" onClick={() => setIsOpen(!isOpen)} />

              }


              <div
                className={`mr-10 col-span-2`}
              >{
                  !isLoading ?
                    <p className={`cursor-pointer font-normal hover:text-blue-800 ${item.service == 'others' && 'hidden'}`}
                      onClick={() => handleNewInstance(item.service)}>
                      Create new instance

                    </p>
                    :
                    <Oval
                      type="Oval"
                      color="#000"
                      width={20}
                      height={20}
                    />
                }

              </div>
            </div>
            <div className={`${isOpen ? '' : 'hidden'}`}>

              <ContainersList containers={item.containers(containers)} refresh={refresh} setRefresh={setRefresh} />
            </div>
          </div>
    )
}

export default ServiceRow