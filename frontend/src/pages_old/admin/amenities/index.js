import AddAmenity from "@/components/admin/AddAmenity"
import AmenitiesList from "@/components/admin/AmenitiesList"
import DashboardAdmin from "@/layouts/DashboardAdmin"
import { UserContext } from "@/layouts/LayoutContext"
import { deleteAmenitie, getAmenities } from "@/lib/api/hotel"
import { useRouter } from "next/router"
import { useContext, useEffect, useState } from "react"

const Dashboard = () => {
  const [user, setUser] = useContext(UserContext);
  const router = useRouter()

  const [amenities, setAmenities] = useState(null);

  const [changed, setChanged] = useState(false);

  const get_amenities = async () => {
    const data = await getAmenities();
    setAmenities(data);
  };

  // useEffect(() => { get_amenities(); }, []);
  
  useEffect(() => { get_amenities(); }, [changed]);

  const handleDeleteAmenitie = async (amenitieID) => {
    //console.log('amenitieID:', amenitieID); // Verificar el valor de amenitieID
    try {
      // Llama a la función deleteAmenitie para eliminar la amenitie
      await deleteAmenitie(amenitieID);

      setChanged(!changed);
    
    } catch (error) {
      console.error('Error al eliminar la comodidad:', error);
    }
  };

  useEffect(() => {

    if (user === null || user.role !== "admin")
      router.push("/auth/login")
  }, [])
  return (
    <DashboardAdmin title={"Amenities List"} current={"/admin/amenities"}>

      {/* Replace with your content */}
      <AddAmenity changed = {changed} setChanged = {setChanged}/>
      <AmenitiesList amenities = {amenities} handleDeleteAmenitie = {handleDeleteAmenitie}/>
      {/* /End replace */}

    </DashboardAdmin>
  )
}

export default Dashboard