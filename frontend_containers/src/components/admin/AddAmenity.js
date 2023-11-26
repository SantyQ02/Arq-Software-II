import { ArrowRightIcon } from "@heroicons/react/solid";
import { insertAmenitie } from "@/lib/api/hotel";
import { useEffect, useState } from "react";
    
export default function AddAmenity({changed, setChanged}) {
  const [formData, setFormData] = useState({
    title: "",
  });

  const { title } = formData;

  const onChange = (e) => setFormData({ ...formData, [e.target.name]: e.target.value });

  const onSubmit = async (e) => {
    e.preventDefault();
    const newAmenitie = await insertAmenitie({ title });
    setChanged(!changed);
  };

  return (
    <div className="">
      <div className="divide-y divide-gray-100">
        <form className="flex justify-start items-end gap-x-6 py-5" onSubmit={onSubmit}>
          <div className="flex gap-x-4">
            <div>
              <label htmlFor="amenity" className="block text-sm font-medium leading-6 text-gray-900">
                Add Amenity
              </label>
              <div className="mt-2">
                <input
                  id="amenity"
                  name="title"
                  value={title}
                  onChange={onChange}
                  placeholder="Amenity Title"
                  type="text"
                  autoComplete="amenity"
                  required
                  className="block w-full rounded-md border-0 p-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                />
              </div>
            </div>
          </div>
          <button type="submit">
            <ArrowRightIcon className="h-6 w-6" aria-hidden="true" />
          </button>
        </form>
      </div>
    </div>
  );
}
