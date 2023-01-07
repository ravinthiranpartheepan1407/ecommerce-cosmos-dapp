package keeper

func (k keeper) GetPostCount(ctx sdk.Context) uint64{
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil{
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k keeper) SetPostCount(ctx sdk.Context, count uint64){
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	byteKey := []byte(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k keeper) AppendPost(ctx sdk.Context, post types.Post) uint64{
	count := k.GetPostCount(ctx)
	post.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)
	appenedValue := k.cdc.MustMarshal(&post)
	store.Set(byteKey, appenedValue)
	k.SetPostCount(ctx, count+1)
	return count
}