import ContainersList from "@/components/admin/ContainersList"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import { getContainersStats, createContainer } from "@/lib/api/containers"
import { useEffect, useState } from "react"
import { PlusIcon, MinusIcon } from '@heroicons/react/solid'

const filtered_containers = [
  {
    service: "frontend",
    title: "Frontend Services",
    containers: (cont) => cont.filter(container => container.name.includes('frontend'))
  },
  {
    service: "hotels",
    title: "Hotels Services",
    containers: (cont) => cont.filter(container => container.name.includes('hotels'))
  },
  {
    service: "search",
    title: "Search Services",
    containers: (cont) => cont.filter(container => container.name.includes('search'))
  },
  {
    service: "business",
    title: "Business Services",
    containers: (cont) => cont.filter(container => container.name.includes('business'))
  },
  {
    service: "others",
    title: "Other Services",
    containers: (cont) => cont.filter(container => ['search', 'frontend', 'business', 'hotels'].every(service => !container.name.includes(service)))
  },
]


const Dashboard = () => {
  const [containers, setContainers] = useState(null);
  const [openSearch, setOpenSearch] = useState(false);
  const [openBusiness, setOpenBusiness] = useState(false);
  const [openOthers, setOpenOthers] = useState(false);
  const [openHotels, setOpenHotels] = useState(false);
  const [openFrontend, setOpenFrontend] = useState(false);
  const [refresh, setRefresh] = useState(false)

  const dic = {
    frontend: {
      setOpen: () => setOpenFrontend(!openFrontend),
      isOpen: () => openFrontend
    },
    search: {
      setOpen: () => setOpenSearch(!openSearch),
      isOpen: () => openSearch
    },
    business: {
      setOpen: () => setOpenBusiness(!openBusiness),
      isOpen: () => openBusiness
    },
    hotels: {
      setOpen: () => setOpenHotels(!openHotels),
      isOpen: () => openHotels
    },
    others: {
      setOpen: () => setOpenOthers(!openOthers),
      isOpen: () => openOthers
    },
  }

  const handleOpen = (service) => {
    dic[service]["setOpen"]()
    console.log(dic[service]["isOpen"]())
  }

  const handleNewInstance = async (service) => {
    await createContainer(service)
    setRefresh(!refresh)
  }

  useEffect(() => {

    const get_containers = async () => {
      const stats = await getContainersStats()
      if (stats != null)
        setContainers(stats)
    }
    get_containers()

  }, [refresh])

  useEffect(() => {
    const fetchData = async () => {
      try {
        const stats = await getContainersStats();
        if (stats != null) {

          setContainers(stats);
        }
      } catch (error) {
        console.error('Error al obtener estadísticas de contenedores:', error);
      }
    };

    // Llama a fetchData inmediatamente y luego cada 10 segundos
    fetchData();

    const intervalId = setInterval(() => {
      fetchData();
    }, 5000);

    // Limpia el intervalo cuando el componente se desmonta
    return () => clearInterval(intervalId);

    // La dependencia refresh aún se mantiene si necesitas reiniciar el intervalo al cambiar refresh
  }, []);

  return (
    <DashboardAdmin title={"Containers Managment"} current={"/admin"}>

      {/* Replace with your content */}
      <div className="divide-y-2 divide-dashed divide-gray-200">
        {containers && filtered_containers.map(item =>
          <div className="pt-5 divide-y-2 divide-dashed">
            <div className="mb-5 font-lg font-bold flex justify-between items-center">
              <div className="flex-1">{item.title}</div>
              <div className={`flex-1 cursor-pointer font-normal hover:text-blue-800 ${item.service == 'others' && 'hidden'}`} onClick={() => handleNewInstance(item.service)}>Create new instance</div>
              {
                dic[item.service]["isOpen"]() ?
                  <MinusIcon className="flex-shrink-0 h-7 w-7 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleOpen(item.service)} />
                  :
                  <PlusIcon className="flex-shrink-0 h-7 w-7 text-blue-400 cursor-pointer" aria-hidden="true" onClick={() => handleOpen(item.service)} />

              }
            </div>
            <div className={`${dic[item.service]["isOpen"]() ? '' : 'hidden'}`}>

              <ContainersList containers={item.containers(containers)} refresh={refresh} setRefresh={setRefresh} />
            </div>
          </div>)}
      </div>
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default Dashboard