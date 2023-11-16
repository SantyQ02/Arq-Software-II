import DashboardAdmin from "@/layouts/DashboardAdmin"
import { getContainersStats } from "@/lib/api/containers"
import { useEffect, useState } from "react"
import ServiceRow from "@/components/admin/ServiceRow"

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
  const [refresh, setRefresh] = useState(false)

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
          <ServiceRow item={item} refresh={refresh} setRefresh={setRefresh} containers={containers} />
          )}
      </div>
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default Dashboard