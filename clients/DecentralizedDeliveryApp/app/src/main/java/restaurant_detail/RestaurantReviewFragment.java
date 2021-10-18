package restaurant_detail;

import android.content.Context;
import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ImageView;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.recyclerview.widget.DividerItemDecoration;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import com.google.android.material.bottomsheet.BottomSheetBehavior;

import java.util.ArrayList;
import java.util.List;

import model.Menu;
import model.Review;
import temporary.variable.android.decentralizeddeliveryapp.R;
import util.Util;

public class RestaurantReviewFragment extends Fragment {
    private RecyclerView mRecyclerView;
    private ReviewRecyclerAdapter mReviewRecyclerAdapter;

    private List<Review> mReviewList;
    private RestaurantDetailActivity mContext;

    RestaurantReviewFragment() {

    }

    @Override
    public void onAttach(@NonNull Context context) {
        super.onAttach(context);
        mContext = (RestaurantDetailActivity) context;
    }

    @Nullable
    @Override
    public View onCreateView(@NonNull LayoutInflater inflater, @Nullable ViewGroup container, @Nullable Bundle savedInstanceState) {
        View v = inflater.inflate(R.layout.fragment_restaurant_review, container, false);
        mRecyclerView = v.findViewById(R.id.recyclerView);

        mReviewList = new ArrayList<>();

        mReviewList.add(new Review("메뉴이름1","맛있어요!!!", R.drawable.chicken_1));
        mReviewList.add(new Review("메뉴이름2","굿", R.drawable.chicken_2));
        mReviewList.add(new Review("메뉴이름3","JMT", R.drawable.chicken_3));
        mReviewList.add(new Review("메뉴이름4","별로입니다. 다시는 안시켜먹어요", R.drawable.chicken_4));

        mReviewRecyclerAdapter = new ReviewRecyclerAdapter(mReviewList);

        mRecyclerView.setLayoutManager(new LinearLayoutManager(getContext()));
        mRecyclerView.addItemDecoration(new DividerItemDecoration(getContext(), LinearLayoutManager.VERTICAL));
        mRecyclerView.setAdapter(mReviewRecyclerAdapter);

        return v;
    }


    class ReviewRecyclerAdapter extends RecyclerView.Adapter<ReviewViewHolder> {
        List<Review> reviewList;

        public ReviewRecyclerAdapter(List<Review> list) {
            reviewList = list;
        }

        @NonNull
        @Override
        public ReviewViewHolder onCreateViewHolder(@NonNull ViewGroup parent, int viewType) {
            View v = LayoutInflater.from(parent.getContext()).inflate(R.layout.list_item_review, parent, false);
            return new ReviewViewHolder(v);
        }

        @Override
        public void onBindViewHolder(@NonNull ReviewViewHolder holder, int i) {
            final Review review = reviewList.get(i);
            holder.title.setText(review.getTitle());
            holder.contents.setText(review.getContents());
            holder.menuImage.setImageResource(review.getImage());
        }

        @Override
        public int getItemCount() {
            return reviewList.size();
        }

        public void setReviewList(List<Review> list) {
            reviewList = list;
            notifyDataSetChanged();
        }
    }

    class ReviewViewHolder extends RecyclerView.ViewHolder {
        TextView title;
        TextView contents;
        ImageView menuImage;


        public ReviewViewHolder(@NonNull View itemView) {
            super(itemView);
            title = itemView.findViewById(R.id.text_review_title);
            contents = itemView.findViewById(R.id.text_review_contents);
            menuImage = itemView.findViewById(R.id.image_review);
        }
    }
}
