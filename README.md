GO LANG REST FUL API LIB

For the PUT method the update operation on the database should change only the fields reported in the call. To do this you must receive the data dynamically in the Map and move it to the structure before sending it to the bank. The FillStruct function transfers values from Map fields to structure.

func restlib.FillStruct(data map[string]interface{}, result interface{}) error

Exemple:

	api.Put("/aluno/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var fields map[string]interface{}
		err := json.Unmarshal(c.Body(), &fields)
		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		aluno := &Aluno{}
		err = mgm.Coll(aluno).FindByID(id, aluno)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		err = restlib.FillStruct(fields, aluno)
		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		err = mgm.Coll(aluno).Update(aluno)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusOK)
	})
