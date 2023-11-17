import { XIcon } from "@heroicons/react/solid";
import { getAmenities, deleteAmenitie } from "@/lib/api/hotel";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function AmenitiesList({amenities, handleDeleteAmenitie}) {
  
  return (
    <div className="">
      <ul role="list" className="divide-y divide-gray-100">
        {amenities !== null && amenities.map((amenitie) => (
          <li key={amenitie._id} className="flex justify-between gap-x-6 py-5">
            <div className="flex gap-x-4">
              <div className="h-12 w-12 flex-none rounded-full bg-gray-50 text-center text-lg font-semibold leading-10 text-blue-700">
                {amenitie.title[0].toUpperCase()}
              </div>
              <div className="min-w-0 flex-auto">
                <p className="text-sm font-semibold leading-10 text-gray-900">{amenitie.title.toUpperCase()}</p>
              </div>
            </div>
            <button type="button" onClick={() => handleDeleteAmenitie(amenitie.amenitie_id)}>
              <XIcon className="h-6 w-6" aria-hidden="true" />
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}
