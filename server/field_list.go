package server

// func (c *Conn) handleFieldList(data []byte) error {
// 	index := bytes.IndexByte(data, 0x00)
// 	table := string(data[0:index])
// 	wildcard := string(data[index+1:])

// 	if c.schema == nil {
// 		return NewDefaultError(ER_NO_DB_ERROR)
// 	}

// 	nodeName := c.schema.rule.GetRule(table).Nodes[0]

// 	n := c.server.getNode(nodeName)

// 	co, err := n.getMasterConn()
// 	if err != nil {
// 		return err
// 	}
// 	defer co.Close()

// 	if err = co.UseDB(c.schema.db); err != nil {
// 		return err
// 	}

// 	if fs, err := co.FieldList(table, wildcard); err != nil {
// 		return err
// 	} else {
// 		return c.writeFieldList(c.status, fs)
// 	}
// }

// func (c *Conn) writeFieldList(status uint16, fs []*Field) error {
// 	c.affectedRows = int64(-1)

// 	data := make([]byte, 4, 1024)

// 	for _, v := range fs {
// 		data = data[0:4]
// 		data = append(data, v.Dump()...)
// 		if err := c.writePacket(data); err != nil {
// 			return err
// 		}
// 	}

// 	if err := c.writeEOF(status); err != nil {
// 		return err
// 	}
// 	return nil
// }
